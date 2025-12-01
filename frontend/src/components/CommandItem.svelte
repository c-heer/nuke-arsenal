<script>
    import Copy from 'phosphor-svelte/lib/Copy';
    import PencilSimple from 'phosphor-svelte/lib/PencilSimple';
    import Trash from 'phosphor-svelte/lib/Trash';
    import CaretDown from 'phosphor-svelte/lib/CaretDown';
    import CaretRight from 'phosphor-svelte/lib/CaretRight';

    let { command, onEdit, onDelete } = $props();
    let expanded = $state(false);
    let copied = $state(false);

    async function copyToClipboard() {
        try {
            await navigator.clipboard.writeText(command.cmd);
            copied = true;
            setTimeout(() => copied = false, 1500);
        } catch (e) {
            console.error('Failed to copy:', e);
        }
    }
</script>

<div class="command-item" class:expanded>
    <div class="command-header" onclick={() => expanded = !expanded}>
        <button class="expand-btn" type="button">
            {#if expanded}
                <CaretDown size={14} />
            {:else}
                <CaretRight size={14} />
            {/if}
        </button>
        <div class="command-main">
            <code class="cmd">{command.cmd}</code>
            <p class="desc">{command.description}</p>
        </div>
        <div class="actions" onclick={(e) => e.stopPropagation()}>
            <button class="action-btn" onclick={copyToClipboard} title="Copy">
                <Copy size={16} />
                {#if copied}<span class="copied">Copied!</span>{/if}
            </button>
            <button class="action-btn" onclick={onEdit} title="Edit">
                <PencilSimple size={16} />
            </button>
            <button class="action-btn delete" onclick={onDelete} title="Delete">
                <Trash size={16} />
            </button>
        </div>
    </div>

    {#if expanded}
        <div class="command-details">
            {#if command.output}
                <div class="detail">
                    <span class="label">Output:</span>
                    <span>{command.output}</span>
                </div>
            {/if}
            {#if command.note}
                <div class="detail">
                    <span class="label">Note:</span>
                    <span>{command.note}</span>
                </div>
            {/if}
            {#if command.tags?.length > 0}
                <div class="tags">
                    {#each command.tags as tag}
                        <span class="tag">{tag}</span>
                    {/each}
                </div>
            {/if}
        </div>
    {/if}
</div>

<style>
    .command-item {
        background: var(--background-2);
        border: 1px solid var(--border-color-1);
        border-radius: var(--radius-1);
        transition: all 0.15s;
    }

    .command-item:hover {
        border-color: var(--border-color-2);
    }

    .command-header {
        display: flex;
        align-items: flex-start;
        gap: var(--space-2);
        padding: var(--space-3);
        cursor: pointer;
    }

    .expand-btn {
        color: var(--on-background-light);
        padding: var(--space-1);
        margin-top: 2px;
    }

    .command-main {
        flex: 1;
        min-width: 0;
    }

    .cmd {
        display: block;
        font-size: var(--font-size-2);
        margin-bottom: var(--space-1);
        word-break: break-all;
    }

    .desc {
        font-size: var(--font-size-1);
        color: var(--on-background-light);
    }

    .actions {
        display: flex;
        gap: var(--space-1);
        opacity: 0;
        transition: opacity 0.15s;
    }

    .command-item:hover .actions {
        opacity: 1;
    }

    .action-btn {
        padding: var(--space-1);
        color: var(--on-background-light);
        border-radius: var(--radius-1);
        position: relative;
    }

    .action-btn:hover {
        color: var(--color-1);
        background: var(--background-3);
    }

    .action-btn.delete:hover {
        color: var(--color-error);
    }

    .copied {
        position: absolute;
        top: -24px;
        left: 50%;
        transform: translateX(-50%);
        font-size: 10px;
        background: var(--color-1);
        color: var(--background-1);
        padding: 2px 6px;
        border-radius: var(--radius-1);
        white-space: nowrap;
    }

    .command-details {
        padding: 0 var(--space-3) var(--space-3);
        padding-left: calc(var(--space-3) + 24px);
        border-top: 1px solid var(--border-color-1);
        margin-top: 0;
    }

    .detail {
        font-size: var(--font-size-1);
        margin-top: var(--space-2);
    }

    .detail .label {
        color: var(--on-background-light);
        margin-right: var(--space-2);
    }

    .tags {
        display: flex;
        gap: var(--space-1);
        margin-top: var(--space-2);
        flex-wrap: wrap;
    }
</style>
