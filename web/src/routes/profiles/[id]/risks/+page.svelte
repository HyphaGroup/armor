<script lang="ts">
	import { page } from '$app/stores';
	import { api } from '$lib/api';

	let loading = $state(true);
	let saving = $state(false);
	let error = $state('');
	let success = $state('');

	let data = $state({
		risks: [] as Array<{
			risk_id: string;
			scenario: string;
			asset_value_score: number;
			likelihood_score: number;
			vulnerability_score: number;
			existing_controls: string;
			control_gaps: string;
			status: string;
		}>,
	});

	const statusOptions = [
		{ value: 'identified', label: 'Identified' },
		{ value: 'accepted', label: 'Accepted' },
		{ value: 'mitigating', label: 'Mitigating' },
		{ value: 'mitigated', label: 'Mitigated' },
	];

	$effect(() => {
		const id = $page.params.id;
		if (id) loadData(id);
	});

	async function loadData(id: string) {
		loading = true;
		error = '';
		try {
			const result = await api.getSection(id, 'risks');
			if (result.data) {
				data = { risks: result.data.risks || [] };
			}
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to load data';
		} finally {
			loading = false;
		}
	}

	async function save() {
		saving = true;
		error = '';
		success = '';
		try {
			await api.updateSection($page.params.id, 'risks', data);
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
				<p class="text-gray-600">Define risk scenarios combining adversaries, threats, and assets.</p>
				<button
					type="button"
					onclick={addRisk}
					class="bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 transition"
				>
					+ Add Risk
				</button>
			</div>

			{#if data.risks.length === 0}
				<p class="text-gray-500 text-center py-8">No risks defined. Click "Add Risk" to begin.</p>
			{:else}
				<div class="space-y-4">
					{#each data.risks as risk, i}
						{@const score = calculateScore(risk)}
						{@const level = getRiskLevel(score)}
						<div class="border rounded-lg p-4 space-y-3">
							<div class="flex justify-between items-start">
								<div class="flex items-center gap-3">
									<span class="px-2 py-1 rounded text-sm font-medium {level.color}">
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
							
							<div>
								<label class="block text-xs text-gray-500 mb-1">Risk Scenario</label>
								<textarea
									bind:value={risk.scenario}
									placeholder="There is a risk that [adversary] could [threat] affecting [asset], which would impact [area]..."
									rows="2"
									class="w-full px-3 py-2 border rounded-lg text-sm"
								></textarea>
							</div>

							<div class="grid grid-cols-3 gap-3">
								<div>
									<label class="block text-xs text-gray-500 mb-1">Asset Value (1-3)</label>
									<input
										type="number"
										min="1"
										max="3"
										bind:value={risk.asset_value_score}
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
								<div>
									<label class="block text-xs text-gray-500 mb-1">Likelihood (1-3)</label>
									<input
										type="number"
										min="1"
										max="3"
										bind:value={risk.likelihood_score}
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
								<div>
									<label class="block text-xs text-gray-500 mb-1">Vulnerability (1-3)</label>
									<input
										type="number"
										min="1"
										max="3"
										bind:value={risk.vulnerability_score}
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
							</div>

							<div class="grid grid-cols-2 gap-3">
								<div>
									<label class="block text-xs text-gray-500 mb-1">Existing Controls</label>
									<textarea
										bind:value={risk.existing_controls}
										rows="2"
										class="w-full px-3 py-2 border rounded-lg text-sm"
									></textarea>
								</div>
								<div>
									<label class="block text-xs text-gray-500 mb-1">Control Gaps</label>
									<textarea
										bind:value={risk.control_gaps}
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
