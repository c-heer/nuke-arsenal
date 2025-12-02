<script>
    import { onMount } from 'svelte';
    import FolderOpen from 'phosphor-svelte/lib/FolderOpen';

    let dataPath = $state('');
    let saved = $state(false);

    async function loadConfig() {
        try {
            const { GetConfig } = await import('../../bindings/github.com/nuke-studios/nuke-arsenal/internal/services/arsenalservice.js');
            const config = await GetConfig();
            dataPath = config?.data_path || '';
        } catch (e) {
            console.error('Failed to load config:', e);
        }
    }

    async function saveConfig() {
        try {
            const { SetDataPath } = await import('../../bindings/github.com/nuke-studios/nuke-arsenal/internal/services/arsenalservice.js');
            await SetDataPath(dataPath);
            saved = true;
            setTimeout(() => saved = false, 2000);
        } catch (e) {
            console.error('Failed to save config:', e);
        }
    }

    onMount(() => {
        loadConfig();
    });
</script>

<div class="page">
    <header class="page-header drag-region">
        <h1>Settings</h1>
    </header>

    <div class="content">
        <div class="setting-card">
            <h2>Data Location</h2>
            <p class="hint">
                Path to your commands.json file. Use a git repo or iCloud folder to sync across machines.
            </p>

            <div class="input-row">
                <FolderOpen size={18} />
                <input
                    type="text"
                    bind:value={dataPath}
                    placeholder="/path/to/.arsenal/commands.json"
                />
            </div>

            <div class="actions">
                <button class="btn btn-primary" onclick={saveConfig}>
                    {saved ? 'Saved!' : 'Save'}
                </button>
            </div>

            <div class="examples">
                <p class="label">Examples:</p>
                <code>~/.arsenal/commands.json</code>
                <code>~/Projects/dotfiles/.arsenal/commands.json</code>
                <code>~/Library/Mobile Documents/com~apple~CloudDocs/.arsenal/commands.json</code>
            </div>
        </div>
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
    }

    .page-header h1 {
        font-size: var(--font-size-3);
        font-weight: var(--font-weight-medium);
    }

    .content {
        flex: 1;
        padding: var(--space-4);
        overflow-y: auto;
    }

    .setting-card {
        background: var(--background-2);
        border: 1px solid var(--border-color-1);
        border-radius: var(--radius-2);
        padding: var(--space-4);
        max-width: 600px;
    }

    .setting-card h2 {
        font-size: var(--font-size-2);
        margin-bottom: var(--space-2);
    }

    .hint {
        font-size: var(--font-size-1);
        color: var(--on-background-light);
        margin-bottom: var(--space-3);
    }

    .input-row {
        display: flex;
        align-items: center;
        gap: var(--space-2);
        background: var(--background-1);
        border: 1px solid var(--border-color-1);
        border-radius: var(--radius-1);
        padding: var(--space-2) var(--space-3);
    }

    .input-row input {
        flex: 1;
        border: none;
        background: transparent;
    }

    .input-row :global(svg) {
        color: var(--on-background-light);
    }

    .actions {
        margin-top: var(--space-3);
    }

    .examples {
        margin-top: var(--space-4);
        padding-top: var(--space-3);
        border-top: 1px solid var(--border-color-1);
    }

    .examples .label {
        font-size: var(--font-size-1);
        color: var(--on-background-light);
        margin-bottom: var(--space-2);
    }

    .examples code {
        display: block;
        margin-bottom: var(--space-1);
        font-size: var(--font-size-1);
    }
</style>
