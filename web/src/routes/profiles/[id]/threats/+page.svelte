<script lang="ts">
	import { page } from '$app/stores';
	import { api } from '$lib/api';

	let loading = $state(true);
	let saving = $state(false);
	let error = $state('');
	let success = $state('');

	let data = $state({
		threats: [] as Array<{
			threat_id: string;
			name: string;
			category: string;
			description: string;
			likelihood: string;
			example: string;
		}>,
	});

	const categoryOptions = [
		{ value: 'account_access', label: 'Account & Access' },
		{ value: 'data_information', label: 'Data & Information' },
		{ value: 'disruption', label: 'Disruption' },
		{ value: 'information_reputation', label: 'Information & Reputation' },
		{ value: 'physical', label: 'Physical' },
		{ value: 'operational', label: 'Operational' },
	];

	const likelihoodOptions = [
		{ value: 'high', label: 'High (Expected/Active)' },
		{ value: 'medium', label: 'Medium (Possible)' },
		{ value: 'low', label: 'Low (Unlikely)' },
	];

	$effect(() => {
		const id = $page.params.id;
		if (id) loadData(id);
	});

	async function loadData(id: string) {
		loading = true;
		error = '';
		try {
			const result = await api.getSection(id, 'threats');
			if (result.data) {
				data = { threats: result.data.threats || [] };
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
			await api.updateSection(id, 'threats', data);
			success = 'Saved successfully';
			setTimeout(() => success = '', 3000);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to save';
		} finally {
			saving = false;
		}
	}

	function addThreat() {
		data.threats = [...data.threats, {
			threat_id: `threat-${Date.now()}`,
			name: '',
			category: '',
			description: '',
			likelihood: 'medium',
			example: '',
		}];
	}

	function removeThreat(index: number) {
		data.threats = data.threats.filter((_, i) => i !== index);
	}
</script>

<div class="space-y-6">
	<div class="flex justify-between items-center">
		<div>
			<a href="/profiles/{$page.params.id}" class="text-blue-600 hover:text-blue-700 text-sm">
				← Back to profile
			</a>
			<h1 class="text-2xl font-bold mt-2">Threats</h1>
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
				<p class="text-gray-600">Map specific threats relevant to your organization.</p>
				<button
					type="button"
					onclick={addThreat}
					class="bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 transition"
				>
					+ Add Threat
				</button>
			</div>

			{#if data.threats.length === 0}
				<p class="text-gray-500 text-center py-8">No threats defined. Click "Add Threat" to begin.</p>
			{:else}
				<div class="space-y-4">
					{#each data.threats as threat, i}
						<div class="border rounded-lg p-4 space-y-3">
							<div class="flex justify-between items-start">
								<div class="flex-1 grid grid-cols-3 gap-3">
									<div>
										<label class="block text-xs text-gray-500 mb-1">Name</label>
										<input
											type="text"
											bind:value={threat.name}
											placeholder="e.g., Phishing"
											class="w-full px-3 py-2 border rounded-lg text-sm"
										/>
									</div>
									<div>
										<label class="block text-xs text-gray-500 mb-1">Category</label>
										<select bind:value={threat.category} class="w-full px-3 py-2 border rounded-lg text-sm">
											<option value="">Select...</option>
											{#each categoryOptions as opt}
												<option value={opt.value}>{opt.label}</option>
											{/each}
										</select>
									</div>
									<div>
										<label class="block text-xs text-gray-500 mb-1">Likelihood</label>
										<select bind:value={threat.likelihood} class="w-full px-3 py-2 border rounded-lg text-sm">
											{#each likelihoodOptions as opt}
												<option value={opt.value}>{opt.label}</option>
											{/each}
										</select>
									</div>
								</div>
								<button
									type="button"
									onclick={() => removeThreat(i)}
									class="text-gray-400 hover:text-red-500 ml-2 text-xl"
								>
									×
								</button>
							</div>
							<div>
								<label class="block text-xs text-gray-500 mb-1">Description</label>
								<textarea
									bind:value={threat.description}
									rows="2"
									class="w-full px-3 py-2 border rounded-lg text-sm"
								></textarea>
							</div>
							<div>
								<label class="block text-xs text-gray-500 mb-1">Example</label>
								<input
									type="text"
									bind:value={threat.example}
									placeholder="How might this manifest for your organization?"
									class="w-full px-3 py-2 border rounded-lg text-sm"
								/>
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	{/if}
</div>
