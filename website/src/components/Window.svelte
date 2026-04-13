<script>
	let { children = null, title = "Window", x = 30, y = 30, z = 0, minWidth = 110, minHeight = 110, width = 300, height = 200, onMove = null, onClose = null, resizeRatio = null, hasSpace = true } = $props();
    
    // svelte-ignore state_referenced_locally
    let posX = $state(x), posY = $state(y);
    // svelte-ignore state_referenced_locally
    let sizeX = $state(width), sizeY = $state(height);
   
    let moving = $state(false);
    let resizing = $state(false);

	function move(e) {
        if (moving) {
            posX += e.movementX;
            posY += e.movementY;
            onMove?.(posX, posY)
        }
        if (resizing) {
            sizeX = Math.max(minWidth, sizeX + e.movementX);
            if (resizeRatio) {
                sizeY = sizeX / resizeRatio;
            } else {
                sizeY = Math.max(minHeight, sizeY + e.movementY);
            }
        }        
	}
</script>

<style>
    .draggable {
        user-select: none;
        position: absolute;
    }
    .title-bar {
        cursor: move;
    }
    .resize-handle {
        position: absolute;
        bottom: 0;
        right: 0;
        width: 10px;
        height: 10px;
        cursor: nwse-resize;
    }
</style>

<svelte:window on:mouseup={() => { moving = false; resizing = false; }} on:mousemove={move} />

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="win7 draggable" style="left: {posX}px; top: {posY}px; z-index: {z};">
    <div class="window glass active">
    <div class="title-bar" onmousedown={() => moving = true}>
        <div class="title-bar-text">{title}</div>
        <div class="title-bar-controls">
        <button aria-label="Close" onclick={(e) => onClose?.(e)}></button>
        </div>
    </div>
    <div class="window-body {hasSpace ? 'has-space' : ''}" style="width: {sizeX}px; height: {sizeY}px;">
        {@render children?.()}
    </div>
    <div class="resize-handle" onmousedown={() => resizing = true}></div>
    </div>
</div>