<script lang="ts">
	import { onMount } from 'svelte';
	import { api, type ProfileSummary } from '$lib/api';

	let profiles = $state<ProfileSummary[]>([]);
	let loading = $state(true);
	let error = $state('');
	let showIntro = $state(true);

	onMount(async () => {
		try {
			profiles = await api.listProfiles();
			// Hide intro if user has profiles
			if (profiles.length > 0) {
				showIntro = localStorage.getItem('hideIntro') === 'true';
			}
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to load profiles';
		} finally {
			loading = false;
		}
	});

	function dismissIntro() {
		showIntro = false;
		localStorage.setItem('hideIntro', 'true');
	}

	async function deleteProfile(id: string, name: string) {
		if (!confirm(`Delete profile "${name}"?`)) return;
		try {
			await api.deleteProfile(id);
			profiles = profiles.filter(p => p.id !== id);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to delete profile';
		}
	}
</script>

<div class="space-y-6">
	<!-- Intro Section -->
	{#if showIntro}
		<div class="bg-gradient-to-br from-slate-800 to-slate-900 text-white rounded-xl p-6 shadow-lg relative">
			{#if profiles.length > 0}
				<button 
					onclick={dismissIntro}
					class="absolute top-4 right-4 text-slate-400 hover:text-white"
					title="Dismiss"
				>
					<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
						<path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
					</svg>
				</button>
			{/if}
			<h1 class="text-2xl font-bold mb-3">ARMOR</h1>
			<p class="text-slate-300 text-sm mb-4">Adversary Risk Modeling for Organizational Resilience</p>
			<p class="text-slate-200 text-sm leading-relaxed mb-4">
				A security framework for organizations to systematically identify, assess, and address threats across digital, physical, and information domains.
			</p>
			
			<!-- Threat Modeling Chain -->
			<div class="bg-slate-700/50 rounded-lg p-4 mb-4">
				<div class="text-xs text-slate-400 uppercase tracking-wide mb-2">The Threat Modeling Chain</div>
				<div class="flex items-center justify-between text-xs text-center gap-1">
					<div class="flex-1">
						<div class="bg-blue-600 rounded px-2 py-1 font-medium">Mission</div>
						<div class="text-slate-400 mt-1">Why you exist</div>
					</div>
					<div class="text-slate-500">→</div>
					<div class="flex-1">
						<div class="bg-blue-600 rounded px-2 py-1 font-medium">Assets</div>
						<div class="text-slate-400 mt-1">What to protect</div>
					</div>
					<div class="text-slate-500">→</div>
					<div class="flex-1">
						<div class="bg-blue-600 rounded px-2 py-1 font-medium">Adversaries</div>
						<div class="text-slate-400 mt-1">Who targets you</div>
					</div>
					<div class="text-slate-500">→</div>
					<div class="flex-1">
						<div class="bg-blue-600 rounded px-2 py-1 font-medium">Threats</div>
						<div class="text-slate-400 mt-1">What they do</div>
					</div>
					<div class="text-slate-500">→</div>
					<div class="flex-1">
						<div class="bg-blue-600 rounded px-2 py-1 font-medium">Risks</div>
						<div class="text-slate-400 mt-1">What matters</div>
					</div>
					<div class="text-slate-500">→</div>
					<div class="flex-1">
						<div class="bg-blue-600 rounded px-2 py-1 font-medium">Mitigations</div>
						<div class="text-slate-400 mt-1">What to do</div>
					</div>
				</div>
			</div>

			<div class="grid grid-cols-3 gap-3 text-xs">
				<div class="bg-slate-700/50 rounded p-3">
					<div class="font-medium text-blue-400 mb-1">Mission-First</div>
					<div class="text-slate-400">Security protects your mission, not the other way around</div>
				</div>
				<div class="bg-slate-700/50 rounded p-3">
					<div class="font-medium text-blue-400 mb-1">Adversary-Informed</div>
					<div class="text-slate-400">Focus on threats YOUR adversaries actually use</div>
				</div>
				<div class="bg-slate-700/50 rounded p-3">
					<div class="font-medium text-blue-400 mb-1">Proportional Response</div>
					<div class="text-slate-400">Match investment to actual threat levels</div>
				</div>
			</div>
		</div>
	{/if}

	<div class="flex justify-between items-center">
		<h1 class="text-2xl font-bold">Profiles</h1>
		<a
			href="/profiles/new"
			class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
		>
			New Profile
		</a>
	</div>

	{#if loading}
		<p class="text-gray-500">Loading...</p>
	{:else if error}
		<p class="text-red-500">{error}</p>
	{:else if profiles.length === 0}
		<div class="text-center py-12 bg-white rounded-lg shadow">
			<p class="text-gray-500 mb-4">No profiles yet</p>
			<a
				href="/profiles/new"
				class="text-blue-600 hover:text-blue-700"
			>
				Create your first profile
			</a>
		</div>
	{:else}
		<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
			{#each profiles as profile}
				<div class="bg-white rounded-lg shadow p-6 hover:shadow-md transition">
					<div class="flex justify-between items-start mb-4">
						<div>
							<h2 class="font-semibold text-lg">{profile.name}</h2>
							{#if profile.description}
								<p class="text-gray-500 text-sm mt-1">{profile.description}</p>
							{/if}
						</div>
						<button
							onclick={() => deleteProfile(profile.id, profile.name)}
							class="text-gray-400 hover:text-red-500 transition"
							title="Delete profile"
						>
							<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
								<path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
							</svg>
						</button>
					</div>
					
					<div class="mb-4">
						<div class="flex justify-between text-sm mb-1">
							<span class="text-gray-600">Completeness</span>
							<span class="font-medium">{Math.round(profile.completeness)}%</span>
						</div>
						<div class="w-full bg-gray-200 rounded-full h-2">
							<div
								class="bg-blue-600 h-2 rounded-full transition-all"
								style="width: {profile.completeness}%"
							></div>
						</div>
					</div>

					<a
						href="/profiles/{profile.id}"
						class="block text-center bg-gray-100 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-200 transition"
					>
						Open
					</a>
				</div>
			{/each}
		</div>
	{/if}
</div>
