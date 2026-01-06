<script lang="ts">
	import { page } from '$app/stores';
	import { api } from '$lib/api';

	let loading = $state(true);
	let saving = $state(false);
	let error = $state('');
	let success = $state('');

	let data = $state({
		adversaries: [] as Array<{
			adversary_id: string;
			name: string;
			category: string;
			why_target_you: string;
			what_they_want: string;
			relevance: string;
			relevance_rationale: string;
		}>,
	});

	const categoryOptions = [
		{ value: 'nation_state', label: 'Nation-State / Intelligence' },
		{ value: 'ideological_opposition', label: 'Ideological Opposition' },
		{ value: 'cybercriminal', label: 'Cybercriminal' },
		{ value: 'insider', label: 'Insider Threat' },
		{ value: 'competitor', label: 'Competitor / Opposing Org' },
		{ value: 'opportunistic', label: 'Opportunistic' },
	];

	const relevanceOptions = [
		{ value: 'confirmed', label: 'Confirmed' },
		{ value: 'likely', label: 'Likely' },
		{ value: 'possible', label: 'Possible' },
		{ value: 'unlikely', label: 'Unlikely' },
	];

	$effect(() => {
		const id = $page.params.id;
		if (id) loadData(id);
	});

	async function loadData(id: string) {
		loading = true;
		error = '';
		try {
			const result = await api.getSection(id, 'adversaries');
			if (result.data) {
				data = { adversaries: result.data.adversaries || [] };
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
			await api.updateSection($page.params.id, 'adversaries', data);
			success = 'Saved successfully';
			setTimeout(() => success = '', 3000);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to save';
		} finally {
			saving = false;
		}
	}

	function addAdversary() {
		data.adversaries = [...data.adversaries, {
			adversary_id: `adv-${Date.now()}`,
			name: '',
			category: '',
			why_target_you: '',
			what_they_want: '',
			relevance: 'possible',
			relevance_rationale: '',
		}];
	}

	function removeAdversary(index: number) {
		data.adversaries = data.adversaries.filter((_, i) => i !== index);
	}
</script>

<div class="space-y-6">
	<div class="flex justify-between items-center">
		<div>
			<a href="/profiles/{$page.params.id}" class="text-blue-600 hover:text-blue-700 text-sm">
				← Back to profile
			</a>
			<h1 class="text-2xl font-bold mt-2">Adversaries</h1>
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
				<p class="text-gray-600">Identify who might target your organization and why.</p>
				<button
					type="button"
					onclick={addAdversary}
					class="bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 transition"
				>
					+ Add Adversary
				</button>
			</div>

			{#if data.adversaries.length === 0}
				<p class="text-gray-500 text-center py-8">No adversaries defined. Click "Add Adversary" to begin.</p>
			{:else}
				<div class="space-y-4">
					{#each data.adversaries as adv, i}
						<div class="border rounded-lg p-4 space-y-3">
							<div class="flex justify-between items-start">
								<div class="flex-1 grid grid-cols-2 gap-3">
									<div>
										<label class="block text-xs text-gray-500 mb-1">Name/Label</label>
										<input
											type="text"
											bind:value={adv.name}
											placeholder="e.g., State Actor X"
											class="w-full px-3 py-2 border rounded-lg text-sm"
										/>
									</div>
									<div>
										<label class="block text-xs text-gray-500 mb-1">Category</label>
										<select bind:value={adv.category} class="w-full px-3 py-2 border rounded-lg text-sm">
											<option value="">Select...</option>
											{#each categoryOptions as opt}
												<option value={opt.value}>{opt.label}</option>
											{/each}
										</select>
									</div>
								</div>
								<button
									type="button"
									onclick={() => removeAdversary(i)}
									class="text-gray-400 hover:text-red-500 ml-2 text-xl"
								>
									×
								</button>
							</div>
							<div class="grid grid-cols-2 gap-3">
								<div>
									<label class="block text-xs text-gray-500 mb-1">Why would they target you?</label>
									<textarea
										bind:value={adv.why_target_you}
										rows="2"
										class="w-full px-3 py-2 border rounded-lg text-sm"
									></textarea>
								</div>
								<div>
									<label class="block text-xs text-gray-500 mb-1">What do they want?</label>
									<textarea
										bind:value={adv.what_they_want}
										rows="2"
										class="w-full px-3 py-2 border rounded-lg text-sm"
									></textarea>
								</div>
							</div>
							<div class="grid grid-cols-2 gap-3">
								<div>
									<label class="block text-xs text-gray-500 mb-1">Relevance</label>
									<select bind:value={adv.relevance} class="w-full px-3 py-2 border rounded-lg text-sm">
										{#each relevanceOptions as opt}
											<option value={opt.value}>{opt.label}</option>
										{/each}
									</select>
								</div>
								<div>
									<label class="block text-xs text-gray-500 mb-1">Relevance Rationale</label>
									<input
										type="text"
										bind:value={adv.relevance_rationale}
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
