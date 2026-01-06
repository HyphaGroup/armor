<script lang="ts">
	import { page } from '$app/stores';
	import { api } from '$lib/api';

	let loading = $state(true);
	let saving = $state(false);
	let error = $state('');
	let success = $state('');

	let data = $state({
		mitigations: [] as Array<{
			mitigation_id: string;
			title: string;
			description: string;
			type: string;
			priority: string;
			effort: string;
			owner: string;
			status: string;
			target_date: string;
			addresses_risks: string[];
		}>,
	});

	// Related data for dropdowns
	let risks = $state<Array<{ risk_id: string; scenario: string; status: string }>>([]);

	const typeOptions = [
		{ value: 'technical', label: 'Technical Control' },
		{ value: 'policy', label: 'Policy/Procedure' },
		{ value: 'training', label: 'Training/Awareness' },
		{ value: 'physical', label: 'Physical Control' },
		{ value: 'organizational', label: 'Organizational' },
	];

	const priorityOptions = [
		{ value: 'critical', label: 'Critical' },
		{ value: 'high', label: 'High' },
		{ value: 'medium', label: 'Medium' },
		{ value: 'low', label: 'Low' },
	];

	const effortOptions = [
		{ value: 'minimal', label: 'Minimal' },
		{ value: 'low', label: 'Low' },
		{ value: 'medium', label: 'Medium' },
		{ value: 'high', label: 'High' },
		{ value: 'major', label: 'Major' },
	];

	const statusOptions = [
		{ value: 'planned', label: 'Planned' },
		{ value: 'in_progress', label: 'In Progress' },
		{ value: 'completed', label: 'Completed' },
		{ value: 'deferred', label: 'Deferred' },
	];

	$effect(() => {
		const id = $page.params.id;
		if (id) loadAllData(id);
	});

	async function loadAllData(id: string) {
		loading = true;
		error = '';
		try {
			const [mitigationsRes, risksRes] = await Promise.all([
				api.getSection(id, 'mitigations'),
				api.getSection(id, 'risks'),
			]);
			
			if (mitigationsRes.data) {
				// Ensure addresses_risks array exists on all mitigations
				const mitigations = (mitigationsRes.data.mitigations || []).map((m: any) => ({
					...m,
					addresses_risks: m.addresses_risks || [],
				}));
				data = { mitigations };
			}
			if (risksRes.data) {
				risks = risksRes.data.risks || [];
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
			await api.updateSection($page.params.id, 'mitigations', data);
			success = 'Saved successfully';
			setTimeout(() => success = '', 3000);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to save';
		} finally {
			saving = false;
		}
	}

	function addMitigation() {
		data.mitigations = [...data.mitigations, {
			mitigation_id: `mit-${Date.now()}`,
			title: '',
			description: '',
			type: 'technical',
			priority: 'medium',
			effort: 'medium',
			owner: '',
			status: 'planned',
			target_date: '',
			addresses_risks: [],
		}];
	}

	function removeMitigation(index: number) {
		data.mitigations = data.mitigations.filter((_, i) => i !== index);
	}

	function toggleRisk(mitIndex: number, riskId: string) {
		const mit = data.mitigations[mitIndex];
		if (mit.addresses_risks.includes(riskId)) {
			data.mitigations[mitIndex].addresses_risks = mit.addresses_risks.filter(id => id !== riskId);
		} else {
			data.mitigations[mitIndex].addresses_risks = [...mit.addresses_risks, riskId];
		}
	}

	function getRiskLabel(risk: typeof risks[0]): string {
		const scenario = risk.scenario || '';
		return scenario.length > 60 ? scenario.substring(0, 60) + '...' : scenario || risk.risk_id;
	}
</script>

<div class="space-y-6">
	<div class="flex justify-between items-center">
		<div>
			<a href="/profiles/{$page.params.id}" class="text-blue-600 hover:text-blue-700 text-sm">
				← Back to profile
			</a>
			<h1 class="text-2xl font-bold mt-2">Mitigations</h1>
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
				<p class="text-gray-600">Plan mitigations to address identified risks.</p>
				<button
					type="button"
					onclick={addMitigation}
					class="bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 transition"
				>
					+ Add Mitigation
				</button>
			</div>

			{#if risks.length === 0}
				<div class="bg-yellow-50 border border-yellow-200 text-yellow-800 px-4 py-3 rounded-lg mb-4">
					<strong>Tip:</strong> Define your <a href="/profiles/{$page.params.id}/risks" class="underline">Risks</a> first to link mitigations to specific risks.
				</div>
			{/if}

			{#if data.mitigations.length === 0}
				<p class="text-gray-500 text-center py-8">No mitigations defined. Click "Add Mitigation" to begin.</p>
			{:else}
				<div class="space-y-6">
					{#each data.mitigations as mit, i}
						<div class="border rounded-lg p-4 space-y-4">
							<div class="flex justify-between items-start">
								<div class="flex-1">
									<label for="title-{i}" class="block text-xs text-gray-500 mb-1">Title</label>
									<input
										id="title-{i}"
										type="text"
										bind:value={mit.title}
										placeholder="e.g., Enable MFA on all accounts"
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
								<button
									type="button"
									onclick={() => removeMitigation(i)}
									class="text-gray-400 hover:text-red-500 ml-2 text-xl"
								>
									×
								</button>
							</div>

							<div>
								<label for="desc-{i}" class="block text-xs text-gray-500 mb-1">Description</label>
								<textarea
									id="desc-{i}"
									bind:value={mit.description}
									rows="2"
									class="w-full px-3 py-2 border rounded-lg text-sm"
								></textarea>
							</div>

							<!-- Link to Risks -->
							{#if risks.length > 0}
								<div class="bg-gray-50 rounded-lg p-4">
									<div class="text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
										Addresses Risks ({mit.addresses_risks.length} selected)
									</div>
									<div class="flex flex-wrap gap-2">
										{#each risks as risk}
											<button
												type="button"
												onclick={() => toggleRisk(i, risk.risk_id)}
												class="px-3 py-1 text-xs rounded-full border transition {mit.addresses_risks.includes(risk.risk_id) ? 'bg-blue-100 border-blue-300 text-blue-700' : 'bg-white border-gray-300 text-gray-600 hover:border-gray-400'}"
												title={risk.scenario}
											>
												{getRiskLabel(risk)}
											</button>
										{/each}
									</div>
								</div>
							{/if}

							<div class="grid grid-cols-4 gap-3">
								<div>
									<label for="type-{i}" class="block text-xs text-gray-500 mb-1">Type</label>
									<select id="type-{i}" bind:value={mit.type} class="w-full px-3 py-2 border rounded-lg text-sm">
										{#each typeOptions as opt}
											<option value={opt.value}>{opt.label}</option>
										{/each}
									</select>
								</div>
								<div>
									<label for="priority-{i}" class="block text-xs text-gray-500 mb-1">Priority</label>
									<select id="priority-{i}" bind:value={mit.priority} class="w-full px-3 py-2 border rounded-lg text-sm">
										{#each priorityOptions as opt}
											<option value={opt.value}>{opt.label}</option>
										{/each}
									</select>
								</div>
								<div>
									<label for="effort-{i}" class="block text-xs text-gray-500 mb-1">Effort</label>
									<select id="effort-{i}" bind:value={mit.effort} class="w-full px-3 py-2 border rounded-lg text-sm">
										{#each effortOptions as opt}
											<option value={opt.value}>{opt.label}</option>
										{/each}
									</select>
								</div>
								<div>
									<label for="status-{i}" class="block text-xs text-gray-500 mb-1">Status</label>
									<select id="status-{i}" bind:value={mit.status} class="w-full px-3 py-2 border rounded-lg text-sm">
										{#each statusOptions as opt}
											<option value={opt.value}>{opt.label}</option>
										{/each}
									</select>
								</div>
							</div>

							<div class="grid grid-cols-2 gap-3">
								<div>
									<label for="owner-{i}" class="block text-xs text-gray-500 mb-1">Owner</label>
									<input
										id="owner-{i}"
										type="text"
										bind:value={mit.owner}
										placeholder="Who is responsible?"
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
								<div>
									<label for="date-{i}" class="block text-xs text-gray-500 mb-1">Target Date</label>
									<input
										id="date-{i}"
										type="date"
										bind:value={mit.target_date}
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	{/if}
</div>
