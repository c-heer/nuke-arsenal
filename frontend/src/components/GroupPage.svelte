<script>
    import { onMount } from 'svelte';
    import Plus from 'phosphor-svelte/lib/Plus';
    import CommandItem from './CommandItem.svelte';

    let { groupKey } = $props();
    let group = $state(null);
    let showForm = $state(false);
    let editingCommand = $state(null);

    // Form fields
    let formCmd = $state('');
    let formDesc = $state('');
    let formOutput = $state('');
    let formNote = $state('');
    let formTags = $state('');

    async function loadGroup() {
        try {
            const { GetGroups } = await import('../../bindings/github.com/c-heer/nuke-arsenal/internal/services/arsenalservice.js');
            const groups = await GetGroups();
            group = groups?.[groupKey] || null;
        } catch (e) {
            console.error('Failed to load group:', e);
        }
    }

    function openAddForm() {
        editingCommand = null;
        formCmd = '';
        formDesc = '';
        formOutput = '';
        formNote = '';
        formTags = '';
        showForm = true;
    }

    function openEditForm(cmd) {
        editingCommand = cmd;
        formCmd = cmd.cmd;
        formDesc = cmd.description;
        formOutput = cmd.output || '';
        formNote = cmd.note || '';
        formTags = (cmd.tags || []).join(', ');
        showForm = true;
    }

    async function saveCommand() {
        const tags = formTags.split(',').map(t => t.trim()).filter(Boolean);
        try {
            if (editingCommand) {
                const { UpdateCommand } = await import('../../bindings/github.com/c-heer/nuke-arsenal/internal/services/arsenalservice.js');
                await UpdateCommand(groupKey, editingCommand.id, formCmd, formDesc, formOutput, formNote, tags);
            } else {
                const { AddCommand } = await import('../../bindings/github.com/c-heer/nuke-arsenal/internal/services/arsenalservice.js');
                await AddCommand(groupKey, formCmd, formDesc, formOutput, formNote, tags);
            }
            showForm = false;
            loadGroup();
        } catch (e) {
            console.error('Failed to save command:', e);
        }
    }

    async function deleteCommand(id) {
        try {
            const { DeleteCommand } = await import('../../bindings/github.com/c-heer/nuke-arsenal/internal/services/arsenalservice.js');
            await DeleteCommand(groupKey, id);
            loadGroup();
        } catch (e) {
            console.error('Failed to delete command:', e);
        }
    }

    onMount(() => {
        loadGroup();
    });

    // Reload when groupKey changes
    $effect(() => {
        groupKey;
        loadGroup();
    });
</script>

<div class="page">
    <header class="page-header drag-region">
        <div class="header-info">
            <h1>{group?.name || 'Loading...'}</h1>
            {#if group?.description}
                <p class="description">{group.description}</p>
            {/if}
        </div>
        <button class="btn btn-primary no-drag" onclick={openAddForm}>
            <Plus size={16} />
            Add Command
        </button>
    </header>

    <div class="content">
        {#if showForm}
            <div class="form-card">
                <h2>{editingCommand ? 'Edit Command' : 'New Command'}</h2>
                <div class="form-field">
                    <label>Command</label>
                    <input type="text" bind:value={formCmd} placeholder="docker exec -it..." />
                </div>
                <div class="form-field">
                    <label>Description</label>
                    <input type="text" bind:value={formDesc} placeholder="What does this do?" />
                </div>
                <div class="form-field">
                    <label>Expected Output (optional)</label>
                    <input type="text" bind:value={formOutput} placeholder="Interactive shell..." />
                </div>
                <div class="form-field">
                    <label>Note (optional)</label>
                    <textarea bind:value={formNote} placeholder="Any tips or context..."></textarea>
                </div>
                <div class="form-field">
                    <label>Tags (comma separated)</label>
                    <input type="text" bind:value={formTags} placeholder="debug, shell, container" />
                </div>
                <div class="form-actions">
                    <button class="btn" onclick={() => showForm = false}>Cancel</button>
                    <button class="btn btn-primary" onclick={saveCommand}>Save</button>
                </div>
            </div>
        {/if}

        {#if group?.commands?.length > 0}
            <div class="commands-list">
                {#each group.commands as cmd}
                    <CommandItem
                        command={cmd}
                        onEdit={() => openEditForm(cmd)}
                        onDelete={() => deleteCommand(cmd.id)}
                    />
                {/each}
            </div>
        {:else if group}
            <p class="empty">No commands yet. Add your first one!</p>
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

    .header-info h1 {
        font-size: var(--font-size-3);
        font-weight: var(--font-weight-medium);
    }

    .description {
        font-size: var(--font-size-1);
        color: var(--on-background-light);
        margin-top: var(--space-1);
    }

    .btn {
        display: flex;
        align-items: center;
        gap: var(--space-1);
    }

    .content {
        flex: 1;
        padding: var(--space-4);
        overflow-y: auto;
    }

    .form-card {
        background: var(--background-2);
        border: 1px solid var(--border-color-1);
        border-radius: var(--radius-2);
        padding: var(--space-4);
        margin-bottom: var(--space-4);
    }

    .form-card h2 {
        font-size: var(--font-size-2);
        margin-bottom: var(--space-3);
    }

    .form-field {
        margin-bottom: var(--space-3);
    }

    .form-field label {
        display: block;
        font-size: var(--font-size-1);
        color: var(--on-background-light);
        margin-bottom: var(--space-1);
    }

    .form-field input,
    .form-field textarea {
        width: 100%;
    }

    .form-field textarea {
        min-height: 80px;
        resize: vertical;
    }

    .form-actions {
        display: flex;
        gap: var(--space-2);
        justify-content: flex-end;
    }

    .commands-list {
        display: flex;
        flex-direction: column;
        gap: var(--space-2);
    }

    .empty {
        color: var(--on-background-light);
        text-align: center;
        padding: var(--space-5);
    }
</style>
