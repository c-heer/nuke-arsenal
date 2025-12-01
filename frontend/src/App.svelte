<script>
    import { onMount } from "svelte";
    import NavRail from "./components/NavRail.svelte";
    import Dashboard from "./components/Dashboard.svelte";
    import GroupPage from "./components/GroupPage.svelte";
    import Settings from "./components/Settings.svelte";

    let currentPage = $state('dashboard');
    let currentGroup = $state(null);

    function navigate(page, groupKey = null) {
        currentPage = page;
        currentGroup = groupKey;
    }
</script>

<div class="app-layout">
    <NavRail {currentPage} {currentGroup} {navigate} />

    <main class="main-content">
        {#if currentPage === 'dashboard'}
            <Dashboard {navigate} />
        {:else if currentPage === 'group' && currentGroup}
            <GroupPage groupKey={currentGroup} />
        {:else if currentPage === 'settings'}
            <Settings />
        {/if}
    </main>
</div>

<style>
    .app-layout {
        display: grid;
        grid-template-columns: auto 1fr;
        height: 100vh;
        width: 100vw;
        overflow: hidden;
    }

    .main-content {
        overflow-y: auto;
        overflow-x: hidden;
    }
</style>
