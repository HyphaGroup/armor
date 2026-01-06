<script lang="ts">
	import { page } from '$app/stores';
	import { api } from '$lib/api';

	let loading = $state(true);
	let saving = $state(false);
	let error = $state('');
	let success = $state('');

	let data = $state({
		assets: [] as Array<{
			asset_id: string;
			name: string;
			description: string;
			category: string;
			value: string;
			primary_requirement: string;
		}>,
	});

	const categoryOptions = [
		{ value: 'beneficiary_data', label: 'Beneficiary/Client Data' },
		{ value: 'source_data', label: 'Source Data' },
		{ value: 'donor_supporter_data', label: 'Donor/Supporter Data' },
		{ value: 'staff_volunteer_data', label: 'Staff/Volunteer Data' },
		{ value: 'financial_data', label: 'Financial Data' },
		{ value: 'strategic_data', label: 'Strategic Data' },
		{ value: 'communications', label: 'Communications' },
		{ value: 'credentials', label: 'Credentials' },
		{ value: 'operational_data', label: 'Operational Data' },
		{ value: 'other', label: 'Other' },
	];

	const valueOptions = [
		{ value: 'critical', label: 'Critical' },
		{ value: 'high', label: 'High' },
		{ value: 'medium', label: 'Medium' },
		{ value: 'low', label: 'Low' },
	];

	const requirementOptions = [
		{ value: 'confidentiality', label: 'Confidentiality' },
		{ value: 'integrity', label: 'Integrity' },
		{ value: 'availability', label: 'Availability' },
	];

	$effect(() => {
		const id = $page.params.id;
		if (id) loadData(id);
	});

	async function loadData(id: string) {
		loading = true;
		error = '';
		try {
			const result = await api.getSection(id, 'assets');
			if (result.data) {
				data = { assets: result.data.assets || [] };
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
			await api.updateSection(id, 'assets', data);
			success = 'Saved successfully';
			setTimeout(() => success = '', 3000);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to save';
		} finally {
			saving = false;
		}
	}

	function addAsset() {
		data.assets = [...data.assets, {
			asset_id: `asset-${Date.now()}`,
			name: '',
			description: '',
			category: '',
			value: 'medium',
			primary_requirement: 'confidentiality',
		}];
	}

	function removeAsset(index: number) {
		data.assets = data.assets.filter((_, i) => i !== index);
	}
</script>

<div class="space-y-6">
	<div class="flex justify-between items-center">
		<div>
			<a href="/profiles/{$page.params.id}" class="text-blue-600 hover:text-blue-700 text-sm">
				← Back to profile
			</a>
			<h1 class="text-2xl font-bold mt-2">Assets</h1>
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
				<p class="text-gray-600">Define the information assets that need protection.</p>
				<button
					type="button"
					onclick={addAsset}
					class="bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 transition"
				>
					+ Add Asset
				</button>
			</div>

			{#if data.assets.length === 0}
				<p class="text-gray-500 text-center py-8">No assets defined. Click "Add Asset" to begin.</p>
			{:else}
				<div class="space-y-4">
					{#each data.assets as asset, i}
						<div class="border rounded-lg p-4 space-y-3">
							<div class="flex justify-between items-start">
								<div class="flex-1 grid grid-cols-2 gap-3">
									<div>
										<label class="block text-xs text-gray-500 mb-1">Name</label>
										<input
											type="text"
											bind:value={asset.name}
											placeholder="Asset name"
											class="w-full px-3 py-2 border rounded-lg text-sm"
										/>
									</div>
									<div>
										<label class="block text-xs text-gray-500 mb-1">Category</label>
										<select bind:value={asset.category} class="w-full px-3 py-2 border rounded-lg text-sm">
											<option value="">Select...</option>
											{#each categoryOptions as opt}
												<option value={opt.value}>{opt.label}</option>
											{/each}
										</select>
									</div>
								</div>
								<button
									type="button"
									onclick={() => removeAsset(i)}
									class="text-gray-400 hover:text-red-500 ml-2 text-xl"
								>
									×
								</button>
							</div>
							<div>
								<label class="block text-xs text-gray-500 mb-1">Description</label>
								<textarea
									bind:value={asset.description}
									placeholder="What is this asset and why does it matter?"
									rows="2"
									class="w-full px-3 py-2 border rounded-lg text-sm"
								></textarea>
							</div>
							<div class="grid grid-cols-2 gap-3">
								<div>
									<label class="block text-xs text-gray-500 mb-1">Value</label>
									<select bind:value={asset.value} class="w-full px-3 py-2 border rounded-lg text-sm">
										{#each valueOptions as opt}
											<option value={opt.value}>{opt.label}</option>
										{/each}
									</select>
								</div>
								<div>
									<label class="block text-xs text-gray-500 mb-1">Primary Requirement</label>
									<select bind:value={asset.primary_requirement} class="w-full px-3 py-2 border rounded-lg text-sm">
										{#each requirementOptions as opt}
											<option value={opt.value}>{opt.label}</option>
										{/each}
									</select>
								</div>
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	{/if}
</div>
