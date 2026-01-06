<script lang="ts">
	import { page } from '$app/stores';
	import { api } from '$lib/api';

	let loading = $state(true);
	let saving = $state(false);
	let error = $state('');
	let success = $state('');

	// Data for this section
	let data = $state({
		risks: [] as Array<{
			risk_id: string;
			scenario: string;
			asset_id: string;
			threat_id: string;
			adversary_id: string;
			asset_value_score: number;
			likelihood_score: number;
			vulnerability_score: number;
			existing_controls: string;
			control_gaps: string;
			status: string;
		}>,
	});

	// Related data for dropdowns
	let assets = $state<Array<{ asset_id: string; name: string; value: string }>>([]);
	let threats = $state<Array<{ threat_id: string; name: string; likelihood: string }>>([]);
	let adversaries = $state<Array<{ adversary_id: string; name: string; relevance: string }>>([]);

	const statusOptions = [
		{ value: 'identified', label: 'Identified' },
		{ value: 'accepted', label: 'Accepted' },
		{ value: 'mitigating', label: 'Mitigating' },
		{ value: 'mitigated', label: 'Mitigated' },
	];

	$effect(() => {
		const id = $page.params.id;
		if (id) loadAllData(id);
	});

	async function loadAllData(id: string) {
		loading = true;
		error = '';
		try {
			// Load risks, assets, threats, and adversaries in parallel
			const [risksRes, assetsRes, threatsRes, adversariesRes] = await Promise.all([
				api.getSection(id, 'risks'),
				api.getSection(id, 'assets'),
				api.getSection(id, 'threats'),
				api.getSection(id, 'adversaries'),
			]);
			
			if (risksRes.data) {
				data = { risks: risksRes.data.risks || [] };
			}
			if (assetsRes.data) {
				assets = assetsRes.data.assets || [];
			}
			if (threatsRes.data) {
				threats = threatsRes.data.threats || [];
			}
			if (adversariesRes.data) {
				adversaries = adversariesRes.data.adversaries || [];
			}
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to load data';
		} finally {
			loading = false;
		}
	}

	async function save() {
		const id = $page.params.id;
		if (!id) return;
		saving = true;
		error = '';
		success = '';
		try {
			await api.updateSection(id, 'risks', data);
			success = 'Saved successfully';
			setTimeout(() => success = '', 3000);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to save';
		} finally {
			saving = false;
		}
	}

	function addRisk() {
		data.risks = [...data.risks, {
			risk_id: `risk-${Date.now()}`,
			scenario: '',
			asset_id: '',
			threat_id: '',
			adversary_id: '',
			asset_value_score: 2,
			likelihood_score: 2,
			vulnerability_score: 2,
			existing_controls: '',
			control_gaps: '',
			status: 'identified',
		}];
	}

	function removeRisk(index: number) {
		data.risks = data.risks.filter((_, i) => i !== index);
	}

	function calculateScore(risk: typeof data.risks[0]): number {
		return risk.asset_value_score * risk.likelihood_score * risk.vulnerability_score;
	}

	function getRiskLevel(score: number): { label: string; color: string } {
		if (score >= 18) return { label: 'Critical', color: 'text-red-600 bg-red-100' };
		if (score >= 10) return { label: 'High', color: 'text-orange-600 bg-orange-100' };
		if (score >= 4) return { label: 'Moderate', color: 'text-yellow-600 bg-yellow-100' };
		return { label: 'Low', color: 'text-green-600 bg-green-100' };
	}

	function generateScenario(risk: typeof data.risks[0]): string {
		const asset = assets.find(a => a.asset_id === risk.asset_id);
		const threat = threats.find(t => t.threat_id === risk.threat_id);
		const adversary = adversaries.find(a => a.adversary_id === risk.adversary_id);
		
		if (!asset || !threat) return '';
		
		const adversaryText = adversary ? adversary.name : 'a threat actor';
		return `There is a risk that ${adversaryText} could execute "${threat.name}" affecting "${asset.name}", which would impact the organization.`;
	}

	function autoGenerateScenario(index: number) {
		const risk = data.risks[index];
		const generated = generateScenario(risk);
		if (generated) {
			data.risks[index].scenario = generated;
		}
	}

	// Auto-populate asset value score when asset is selected
	function onAssetChange(index: number) {
		const risk = data.risks[index];
		const asset = assets.find(a => a.asset_id === risk.asset_id);
		if (asset) {
			// Map asset value to score
			const valueMap: Record<string, number> = { critical: 3, high: 2, medium: 2, low: 1 };
			data.risks[index].asset_value_score = valueMap[asset.value] || 2;
		}
	}

	// Auto-populate likelihood score when threat is selected
	function onThreatChange(index: number) {
		const risk = data.risks[index];
		const threat = threats.find(t => t.threat_id === risk.threat_id);
		if (threat) {
			const likelihoodMap: Record<string, number> = { high: 3, medium: 2, low: 1 };
			data.risks[index].likelihood_score = likelihoodMap[threat.likelihood] || 2;
		}
	}
</script>

<div class="space-y-6">
	<div class="flex justify-between items-center">
		<div>
			<a href="/profiles/{$page.params.id}" class="text-blue-600 hover:text-blue-700 text-sm">
				← Back to profile
			</a>
			<h1 class="text-2xl font-bold mt-2">Risks</h1>
		</div>
		<button
			onclick={save}
			disabled={saving || loading}
			class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition disabled:opacity-50"
		>
			{saving ? 'Saving...' : 'Save'}
		</button>
	</div>

	{#if error}
		<div class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg">{error}</div>
	{/if}
	{#if success}
		<div class="bg-green-50 border border-green-200 text-green-700 px-4 py-3 rounded-lg">{success}</div>
	{/if}

	{#if loading}
		<p class="text-gray-500">Loading...</p>
	{:else}
		<div class="bg-white rounded-lg shadow p-6">
			<div class="flex justify-between items-center mb-4">
				<p class="text-gray-600">Build risks by linking assets, threats, and adversaries.</p>
				<button
					type="button"
					onclick={addRisk}
					class="bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 transition"
				>
					+ Add Risk
				</button>
			</div>

			{#if assets.length === 0 || threats.length === 0}
				<div class="bg-yellow-50 border border-yellow-200 text-yellow-800 px-4 py-3 rounded-lg mb-4">
					<strong>Tip:</strong> Define your <a href="/profiles/{$page.params.id}/assets" class="underline">Assets</a> and <a href="/profiles/{$page.params.id}/threats" class="underline">Threats</a> first to use the guided risk builder.
				</div>
			{/if}

			{#if data.risks.length === 0}
				<p class="text-gray-500 text-center py-8">No risks defined. Click "Add Risk" to begin.</p>
			{:else}
				<div class="space-y-6">
					{#each data.risks as risk, i}
						{@const score = calculateScore(risk)}
						{@const level = getRiskLevel(score)}
						<div class="border rounded-lg p-4 space-y-4">
							<div class="flex justify-between items-start">
								<div class="flex items-center gap-3">
									<span class="px-3 py-1 rounded text-sm font-bold {level.color}">
										{level.label} ({score})
									</span>
									<select bind:value={risk.status} class="px-2 py-1 border rounded text-sm">
										{#each statusOptions as opt}
											<option value={opt.value}>{opt.label}</option>
										{/each}
									</select>
								</div>
								<button
									type="button"
									onclick={() => removeRisk(i)}
									class="text-gray-400 hover:text-red-500 text-xl"
								>
									×
								</button>
							</div>

							<!-- Risk Builder Section -->
							<div class="bg-gray-50 rounded-lg p-4 space-y-3">
								<div class="text-xs font-medium text-gray-500 uppercase tracking-wide">Risk Components</div>
								<div class="grid grid-cols-3 gap-3">
									<div>
										<label for="asset-{i}" class="block text-xs text-gray-500 mb-1">Asset at Risk</label>
										<select 
											id="asset-{i}"
											bind:value={risk.asset_id} 
											onchange={() => onAssetChange(i)}
											class="w-full px-3 py-2 border rounded-lg text-sm"
										>
											<option value="">Select asset...</option>
											{#each assets as asset}
												<option value={asset.asset_id}>{asset.name}</option>
											{/each}
										</select>
									</div>
									<div>
										<label for="threat-{i}" class="block text-xs text-gray-500 mb-1">Threat</label>
										<select 
											id="threat-{i}"
											bind:value={risk.threat_id}
											onchange={() => onThreatChange(i)}
											class="w-full px-3 py-2 border rounded-lg text-sm"
										>
											<option value="">Select threat...</option>
											{#each threats as threat}
												<option value={threat.threat_id}>{threat.name}</option>
											{/each}
										</select>
									</div>
									<div>
										<label for="adversary-{i}" class="block text-xs text-gray-500 mb-1">Adversary (optional)</label>
										<select 
											id="adversary-{i}"
											bind:value={risk.adversary_id} 
											class="w-full px-3 py-2 border rounded-lg text-sm"
										>
											<option value="">Select adversary...</option>
											{#each adversaries as adv}
												<option value={adv.adversary_id}>{adv.name}</option>
											{/each}
										</select>
									</div>
								</div>
								{#if risk.asset_id && risk.threat_id}
									<button
										type="button"
										onclick={() => autoGenerateScenario(i)}
										class="text-sm text-blue-600 hover:text-blue-700"
									>
										↻ Generate scenario from selections
									</button>
								{/if}
							</div>
							
							<!-- Scenario -->
							<div>
								<label for="scenario-{i}" class="block text-xs text-gray-500 mb-1">Risk Scenario</label>
								<textarea
									id="scenario-{i}"
									bind:value={risk.scenario}
									placeholder="There is a risk that [adversary] could [threat] affecting [asset], which would impact [area]..."
									rows="2"
									class="w-full px-3 py-2 border rounded-lg text-sm"
								></textarea>
							</div>

							<!-- Scoring -->
							<div class="grid grid-cols-3 gap-3">
								<div>
									<label for="asset-value-{i}" class="block text-xs text-gray-500 mb-1">
										Asset Value (1-3)
										{#if risk.asset_id}
											{@const asset = assets.find(a => a.asset_id === risk.asset_id)}
											{#if asset}
												<span class="text-gray-400">· {asset.value}</span>
											{/if}
										{/if}
									</label>
									<input
										id="asset-value-{i}"
										type="number"
										min="1"
										max="3"
										bind:value={risk.asset_value_score}
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
								<div>
									<label for="likelihood-{i}" class="block text-xs text-gray-500 mb-1">
										Likelihood (1-3)
										{#if risk.threat_id}
											{@const threat = threats.find(t => t.threat_id === risk.threat_id)}
											{#if threat}
												<span class="text-gray-400">· {threat.likelihood}</span>
											{/if}
										{/if}
									</label>
									<input
										id="likelihood-{i}"
										type="number"
										min="1"
										max="3"
										bind:value={risk.likelihood_score}
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
								<div>
									<label for="vulnerability-{i}" class="block text-xs text-gray-500 mb-1">Vulnerability (1-3)</label>
									<input
										id="vulnerability-{i}"
										type="number"
										min="1"
										max="3"
										bind:value={risk.vulnerability_score}
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
							</div>

							<!-- Controls -->
							<div class="grid grid-cols-2 gap-3">
								<div>
									<label for="controls-{i}" class="block text-xs text-gray-500 mb-1">Existing Controls</label>
									<textarea
										id="controls-{i}"
										bind:value={risk.existing_controls}
										placeholder="What protections currently exist?"
										rows="2"
										class="w-full px-3 py-2 border rounded-lg text-sm"
									></textarea>
								</div>
								<div>
									<label for="gaps-{i}" class="block text-xs text-gray-500 mb-1">Control Gaps</label>
									<textarea
										id="gaps-{i}"
										bind:value={risk.control_gaps}
										placeholder="What's missing or insufficient?"
										rows="2"
										class="w-full px-3 py-2 border rounded-lg text-sm"
									></textarea>
								</div>
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	{/if}
</div>
