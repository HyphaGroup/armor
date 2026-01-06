<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { api, type Profile } from '$lib/api';

	let profile = $state<Profile | null>(null);
	let loading = $state(true);
	let error = $state('');

	const sections = [
		{ key: 'mission', label: 'Mission & Impact', description: 'Define your mission and impact areas' },
		{ key: 'assets', label: 'Assets', description: 'Identify what needs protection' },
		{ key: 'adversaries', label: 'Adversaries', description: 'Identify who might target you' },
		{ key: 'threats', label: 'Threats', description: 'Map specific threats' },
		{ key: 'risks', label: 'Risks', description: 'Assess and prioritize risks' },
		{ key: 'mitigations', label: 'Mitigations', description: 'Plan improvements' },
	];

	$effect(() => {
		const id = $page.params.id;
		if (id) {
			loadProfile(id);
		}
	});

	async function loadProfile(id: string) {
		loading = true;
		error = '';
		try {
			profile = await api.getProfile(id);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to load profile';
		} finally {
			loading = false;
		}
	}

	function getSectionCompleteness(sectionKey: string): number {
		if (!profile?.completeness?.sections) return 0;
		const section = profile.completeness.sections.find(s => s.section === sectionKey);
		return section?.percentage ?? 0;
	}
</script>

{#if loading}
	<p class="text-gray-500">Loading...</p>
{:else if error}
	<p class="text-red-500">{error}</p>
{:else if profile}
	<div class="space-y-6">
		<div class="flex justify-between items-start">
			<div>
				<a href="/profiles" class="text-blue-600 hover:text-blue-700 text-sm mb-2 inline-block">
					← Back to profiles
				</a>
				<h1 class="text-2xl font-bold">{profile.name}</h1>
				{#if profile.description}
					<p class="text-gray-500 mt-1">{profile.description}</p>
				{/if}
			</div>
			<div class="text-right">
				<div class="text-sm text-gray-500">Overall Completeness</div>
				<div class="text-2xl font-bold text-blue-600">{Math.round(profile.completeness.overall)}%</div>
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
			{#each sections as section}
				{@const completeness = getSectionCompleteness(section.key)}
				<a
					href="/profiles/{profile.id}/{section.key}"
					class="bg-white rounded-lg shadow p-6 hover:shadow-md transition block"
				>
					<h2 class="font-semibold text-lg mb-1">{section.label}</h2>
					<p class="text-gray-500 text-sm mb-4">{section.description}</p>
					
					<div class="flex justify-between text-sm mb-1">
						<span class="text-gray-600">Progress</span>
						<span class="font-medium">{Math.round(completeness)}%</span>
					</div>
					<div class="w-full bg-gray-200 rounded-full h-2">
						<div
							class="h-2 rounded-full transition-all"
							class:bg-green-500={completeness === 100}
							class:bg-blue-600={completeness > 0 && completeness < 100}
							class:bg-gray-300={completeness === 0}
							style="width: {completeness}%"
						></div>
					</div>
				</a>
			{/each}
		</div>
	</div>
{/if}
