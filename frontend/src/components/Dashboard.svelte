<script>
    import { onMount } from 'svelte';
    import MagnifyingGlass from 'phosphor-svelte/lib/MagnifyingGlass';

    let { navigate } = $props();
    let groups = $state({});
    let searchQuery = $state('');
    let searchResults = $state([]);
    let totalCommands = $derived(
        Object.values(groups).reduce((sum, g) => sum + (g.commands?.length || 0), 0)
    );

    async function loadGroups() {
        try {
            const { GetGroups } = await import('../../bindings/github.com/nuke-studios/nuke-arsenal/internal/services/arsenalservice.js');
            groups = await GetGroups() || {};
        } catch (e) {
            console.error('Failed to load groups:', e);
        }
    }

    async function search() {
        if (!searchQuery.trim()) {
            searchResults = [];
            return;
        }
        try {
            const { Search } = await import('../../bindings/github.com/nuke-studios/nuke-arsenal/internal/services/arsenalservice.js');
            searchResults = await Search(searchQuery) || [];
        } catch (e) {
            console.error('Search failed:', e);
        }
    }

    onMount(() => {
        loadGroups();
    });
</script>

<div class="page">
    <header class="page-header drag-region">
        <h1>Dashboard</h1>
        <div class="stats">
            {Object.keys(groups).length} groups � {totalCommands} commands
        </div>
    </header>

    <div class="content">
        <div class="search-section no-drag">
            <div class="search-box">
                <MagnifyingGlass size={18} />
                <input
                    type="text"
                    placeholder="Search commands..."
                    bind:value={searchQuery}
                    oninput={search}
                />
            </div>
        </div>

        {#if searchResults.length > 0}
            <div class="results">
                <h2>Results</h2>
                {#each searchResults as result}
                    <div class="result-item">
                        <code>{result.command.cmd}</code>
                        <p>{result.command.description}</p>
                        <span class="group-badge">{result.groupName}</span>
                    </div>
                {/each}
            </div>
        {:else if searchQuery}
            <p class="no-results">No commands found</p>
        {:else}
            <div class="groups-grid">
                {#each Object.entries(groups) as [key, group]}
                    <button class="group-card" onclick={() => navigate('group', key)}>
                        <h3>{group.name}</h3>
                        <p>{group.description}</p>
                        <span class="count">{group.commands?.length || 0} commands</span>
                    </button>
                {/each}

                {#if Object.keys(groups).length === 0}
                    <p class="empty">No groups yet. Add one from the sidebar.</p>
                {/if}
            </div>
        {/if}
    </div>
</div>

<style>
    .page {
        height: 100%;
        display: flex;
        flex-direction: column;
    }

    .page-header {
        padding: var(--space-3) var(--space-4);
        border-bottom: 1px solid var(--border-color-1);
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .page-header h1 {
        font-size: var(--font-size-3);
        font-weight: var(--font-weight-medium);
    }

    .stats {
        font-size: var(--font-size-1);
        color: var(--on-background-light);
    }

    .content {
        flex: 1;
        padding: var(--space-4);
        overflow-y: auto;
    }

    .search-section {
        margin-bottom: var(--space-4);
    }

    .search-box {
        display: flex;
        align-items: center;
        gap: var(--space-2);
        background: var(--background-2);
        border: 1px solid var(--border-color-1);
        border-radius: var(--radius-2);
        padding: var(--space-2) var(--space-3);
    }

    .search-box input {
        flex: 1;
        border: none;
        background: transparent;
        outline: none;
    }

    .search-box :global(svg) {
        color: var(--on-background-light);
    }

    .groups-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
        gap: var(--space-3);
    }

    .group-card {
        text-align: left;
        padding: var(--space-3);
        background: var(--background-2);
        border: 1px solid var(--border-color-1);
        border-radius: var(--radius-2);
        cursor: pointer;
        transition: all 0.15s;
    }

    .group-card:hover {
        border-color: var(--color-1);
    }

    .group-card h3 {
        font-size: var(--font-size-2);
        margin-bottom: var(--space-1);
    }

    .group-card p {
        font-size: var(--font-size-1);
        color: var(--on-background-light);
        margin-bottom: var(--space-2);
    }

    .group-card .count {
        font-size: var(--font-size-1);
        color: var(--color-2);
    }

    .results h2 {
        font-size: var(--font-size-2);
        margin-bottom: var(--space-3);
        color: var(--on-background-light);
    }

    .result-item {
        padding: var(--space-3);
        background: var(--background-2);
        border: 1px solid var(--border-color-1);
        border-radius: var(--radius-1);
        margin-bottom: var(--space-2);
    }

    .result-item code {
        display: block;
        margin-bottom: var(--space-1);
    }

    .result-item p {
        font-size: var(--font-size-1);
        color: var(--on-background-light);
        margin-bottom: var(--space-2);
    }

    .group-badge {
        font-size: var(--font-size-1);
        padding: 2px var(--space-2);
        background: var(--background-3);
        border-radius: var(--radius-1);
        color: var(--color-2);
    }

    .no-results, .empty {
        color: var(--on-background-light);
        text-align: center;
        padding: var(--space-5);
    }
</style>
