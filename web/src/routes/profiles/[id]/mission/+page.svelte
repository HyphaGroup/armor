<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { api } from '$lib/api';

	let loading = $state(true);
	let saving = $state(false);
	let error = $state('');
	let success = $state('');

	let data = $state({
		mission_statement: '',
		impact_areas: [] as Array<{
			area: string;
			priority: number;
			description: string;
			high_impact_threshold: string;
			medium_impact_threshold: string;
			low_impact_threshold: string;
		}>,
		mission_failure_scenarios: [] as string[],
	});

	const impactAreaOptions = [
		{ value: 'safety_security', label: 'Safety & Security' },
		{ value: 'mission_delivery', label: 'Mission Delivery' },
		{ value: 'trust_reputation', label: 'Trust & Reputation' },
		{ value: 'financial_resources', label: 'Financial Resources' },
		{ value: 'legal_compliance', label: 'Legal & Compliance' },
		{ value: 'partner_relations', label: 'Partner Relations' },
	];

	$effect(() => {
		const id = $page.params.id;
		if (id) loadData(id);
	});

	async function loadData(id: string) {
		loading = true;
		error = '';
		try {
			const result = await api.getSection(id, 'mission');
			if (result.data) {
				data = {
					mission_statement: result.data.mission_statement || '',
					impact_areas: result.data.impact_areas || [],
					mission_failure_scenarios: result.data.mission_failure_scenarios || [],
				};
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
			await api.updateSection($page.params.id, 'mission', data);
			success = 'Saved successfully';
			setTimeout(() => success = '', 3000);
		} catch (e: any) {
			if (e.message?.includes('Validation')) {
				error = 'Validation failed. Please check required fields.';
			} else {
				error = e instanceof Error ? e.message : 'Failed to save';
			}
		} finally {
			saving = false;
		}
	}

	function addImpactArea() {
		data.impact_areas = [...data.impact_areas, {
			area: '',
			priority: data.impact_areas.length + 1,
			description: '',
			high_impact_threshold: '',
			medium_impact_threshold: '',
			low_impact_threshold: '',
		}];
	}

	function removeImpactArea(index: number) {
		data.impact_areas = data.impact_areas.filter((_, i) => i !== index);
	}

	function addFailureScenario() {
		data.mission_failure_scenarios = [...data.mission_failure_scenarios, ''];
	}

	function removeFailureScenario(index: number) {
		data.mission_failure_scenarios = data.mission_failure_scenarios.filter((_, i) => i !== index);
	}
</script>

<div class="space-y-6">
	<div class="flex justify-between items-center">
		<div>
			<a href="/profiles/{$page.params.id}" class="text-blue-600 hover:text-blue-700 text-sm">
				← Back to profile
			</a>
			<h1 class="text-2xl font-bold mt-2">Mission & Impact</h1>
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
		<div class="bg-white rounded-lg shadow p-6 space-y-6">
			<div>
				<label for="mission" class="block text-sm font-medium text-gray-700 mb-1">
					Mission Statement <span class="text-red-500">*</span>
				</label>
				<textarea
					id="mission"
					bind:value={data.mission_statement}
					rows="4"
					placeholder="What is your organization trying to achieve? Who depends on you?"
					class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
				></textarea>
			</div>

			<div>
				<div class="flex justify-between items-center mb-2">
					<label class="block text-sm font-medium text-gray-700">
						Impact Areas <span class="text-red-500">*</span>
					</label>
					<button
						type="button"
						onclick={addImpactArea}
						class="text-blue-600 hover:text-blue-700 text-sm"
					>
						+ Add Impact Area
					</button>
				</div>
				
				{#if data.impact_areas.length === 0}
					<p class="text-gray-500 text-sm">No impact areas defined. Click "Add Impact Area" to begin.</p>
				{:else}
					<div class="space-y-4">
						{#each data.impact_areas as area, i}
							<div class="border rounded-lg p-4 space-y-3">
								<div class="flex justify-between items-start">
									<div class="flex-1 grid grid-cols-2 gap-3">
										<div>
											<label class="block text-xs text-gray-500 mb-1">Area</label>
											<select
												bind:value={area.area}
												class="w-full px-3 py-2 border rounded-lg text-sm"
											>
												<option value="">Select area...</option>
												{#each impactAreaOptions as opt}
													<option value={opt.value}>{opt.label}</option>
												{/each}
											</select>
										</div>
										<div>
											<label class="block text-xs text-gray-500 mb-1">Priority (1=highest)</label>
											<input
												type="number"
												min="1"
												max="6"
												bind:value={area.priority}
												class="w-full px-3 py-2 border rounded-lg text-sm"
											/>
										</div>
									</div>
									<button
										type="button"
										onclick={() => removeImpactArea(i)}
										class="text-gray-400 hover:text-red-500 ml-2"
									>
										×
									</button>
								</div>
								<div>
									<label class="block text-xs text-gray-500 mb-1">Description</label>
									<input
										type="text"
										bind:value={area.description}
										placeholder="What this means for your organization"
										class="w-full px-3 py-2 border rounded-lg text-sm"
									/>
								</div>
								<div class="grid grid-cols-3 gap-2">
									<div>
										<label class="block text-xs text-gray-500 mb-1">High Impact</label>
										<input
											type="text"
											bind:value={area.high_impact_threshold}
											placeholder="High threshold"
											class="w-full px-3 py-2 border rounded-lg text-sm"
										/>
									</div>
									<div>
										<label class="block text-xs text-gray-500 mb-1">Medium Impact</label>
										<input
											type="text"
											bind:value={area.medium_impact_threshold}
											placeholder="Medium threshold"
											class="w-full px-3 py-2 border rounded-lg text-sm"
										/>
									</div>
									<div>
										<label class="block text-xs text-gray-500 mb-1">Low Impact</label>
										<input
											type="text"
											bind:value={area.low_impact_threshold}
											placeholder="Low threshold"
											class="w-full px-3 py-2 border rounded-lg text-sm"
										/>
									</div>
								</div>
							</div>
						{/each}
					</div>
				{/if}
			</div>

			<div>
				<div class="flex justify-between items-center mb-2">
					<label class="block text-sm font-medium text-gray-700">
						Mission Failure Scenarios
					</label>
					<button
						type="button"
						onclick={addFailureScenario}
						class="text-blue-600 hover:text-blue-700 text-sm"
					>
						+ Add Scenario
					</button>
				</div>
				
				{#if data.mission_failure_scenarios.length === 0}
					<p class="text-gray-500 text-sm">No failure scenarios defined.</p>
				{:else}
					<div class="space-y-2">
						{#each data.mission_failure_scenarios as scenario, i}
							<div class="flex gap-2">
								<input
									type="text"
									bind:value={data.mission_failure_scenarios[i]}
									placeholder="What would mission failure look like?"
									class="flex-1 px-3 py-2 border rounded-lg text-sm"
								/>
								<button
									type="button"
									onclick={() => removeFailureScenario(i)}
									class="text-gray-400 hover:text-red-500 px-2"
								>
									×
								</button>
							</div>
						{/each}
					</div>
				{/if}
			</div>
		</div>
	{/if}
</div>
