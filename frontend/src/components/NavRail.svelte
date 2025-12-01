<script>
    import { onMount } from 'svelte';
    import House from 'phosphor-svelte/lib/House';
    import Gear from 'phosphor-svelte/lib/Gear';
    import Plus from 'phosphor-svelte/lib/Plus';
    import CaretRight from 'phosphor-svelte/lib/CaretRight';
    import CaretLeft from 'phosphor-svelte/lib/CaretLeft';
    import Terminal from 'phosphor-svelte/lib/Terminal';

    let { currentPage, currentGroup, navigate } = $props();
    let expanded = $state(false);
    let groups = $state({});

    async function loadGroups() {
        try {
            const { GetGroups } = await import('../../bindings/github.com/c-heer/nuke-arsenal/internal/services/arsenalservice.js');
            groups = await GetGroups() || {};
        } catch (e) {
            console.error('Failed to load groups:', e);
        }
    }

    onMount(() => {
        loadGroups();
    });
</script>

<nav class="nav-rail" class:expanded>
    <button
        class="nav-item nav-logo drag-region"
        onclick={() => expanded = !expanded}
        title="Arsenal"
        type="button"
    >
        <div class="nav-icon-container">
            <Terminal size={24} weight="bold" />
        </div>
        {#if expanded}
            <span class="nav-label logo-text">ARSENAL</span>
        {/if}
    </button>

    <button
        class="nav-item no-drag"
        class:active={currentPage === 'dashboard'}
        onclick={() => navigate('dashboard')}
        title="Dashboard"
        type="button"
    >
        <div class="nav-icon-container">
            <House size={20} />
        </div>
        {#if expanded}<span class="nav-label">Dashboard</span>{/if}
    </button>

    <div class="nav-divider"></div>

    {#each Object.entries(groups) as [key, group]}
        <button
            class="nav-item no-drag"
            class:active={currentPage === 'group' && currentGroup === key}
            onclick={() => navigate('group', key)}
            title={group.name}
            type="button"
        >
            <div class="nav-icon-container">
                <Terminal size={20} />
            </div>
            {#if expanded}<span class="nav-label">{group.name}</span>{/if}
        </button>
    {/each}

    <button
        class="nav-item nav-add no-drag"
        onclick={() => {/* TODO: Add group modal */}}
        title="Add Group"
        type="button"
    >
        <div class="nav-icon-container">
            <Plus size={16} />
        </div>
        {#if expanded}<span class="nav-label">Add Group</span>{/if}
    </button>

    <div class="nav-spacer"></div>

    <button
        class="nav-item no-drag"
        class:active={currentPage === 'settings'}
        onclick={() => navigate('settings')}
        title="Settings"
        type="button"
    >
        <div class="nav-icon-container">
            <Gear size={20} />
        </div>
        {#if expanded}<span class="nav-label">Settings</span>{/if}
    </button>

    <button
        class="nav-item no-drag"
        onclick={() => expanded = !expanded}
        title={expanded ? 'Collapse' : 'Expand'}
        type="button"
    >
        <div class="nav-icon-container">
            {#if expanded}
                <CaretLeft size={16} />
            {:else}
                <CaretRight size={16} />
            {/if}
        </div>
        {#if expanded}<span class="nav-label">Collapse</span>{/if}
    </button>
</nav>

<style>
    .nav-rail {
        display: flex;
        flex-direction: column;
        height: 100%;
        background: var(--background-4);
        border-right: 1px solid var(--border-color-1);
        min-width: 50px;
    }

    .nav-rail.expanded {
        min-width: 200px;
    }

    .nav-item {
        display: flex;
        align-items: center;
        gap: var(--space-2);
        height: 44px;
        padding: 0;
        background: transparent;
        border: none;
        color: var(--on-background-light);
        cursor: pointer;
        transition: all 0.15s;
    }

    .nav-item:hover {
        color: var(--color-1);
        background: var(--background-3);
    }

    .nav-item.active {
        color: var(--color-1);
        background: var(--background-3);
    }

    .nav-item.nav-logo {
        height: 50px;
        color: var(--on-background);
    }

    .nav-item.nav-add {
        color: var(--on-background-light);
        opacity: 0.6;
    }

    .nav-item.nav-add:hover {
        opacity: 1;
    }

    .nav-icon-container {
        width: 50px;
        height: 44px;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-shrink: 0;
    }

    .nav-label {
        font-size: var(--font-size-2);
        font-weight: var(--font-weight-medium);
        white-space: nowrap;
        padding-right: var(--space-3);
    }

    .logo-text {
        font-family: var(--family-2);
        font-weight: var(--font-weight-bold);
        letter-spacing: 0.1em;
    }

    .nav-divider {
        height: 1px;
        background: var(--border-color-1);
        margin: var(--space-2) var(--space-2);
    }

    .nav-spacer {
        flex: 1;
    }
</style>
