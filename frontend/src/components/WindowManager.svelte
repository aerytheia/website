<script>
  import Window from "./Window.svelte";

  let windows = $state([])
  let topZ = $state(0)
  let next = 0;

  function getPosition() {
    let x = 30, y = 30
    while (windows.some(w => w.x === x && w.y === y)) {
      x += 20
      y += 20
    }
    return [x, y]
  }

  export function openWindow(snippet, opts = {}) {
    const [x, y] = getPosition()
    console.log(opts)
    windows.push({ id: next++, snippet, z: ++topZ, x, y, opts })
  }

  function bringToFront(id) {
    windows.find(w => w.id === id).z = ++topZ
  }

  function closeWindow(id) {
    const win = windows.find(w => w.id === id)
    win.onClose?.()
    const i = windows.findIndex(w => w.id === id)
    windows.splice(i, 1)
  }

</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
{#each windows as win (win.id)}
  <div onmousedown={() => bringToFront(win.id)}>
    <Window title={win.opts.title} x={win.x} y={win.y} z={win.z} onClose={() => closeWindow(win.id)} onMove={(x, y) => { win.x = x; win.y = y }} {...win.opts}>
      {@render win.snippet()}
    </Window>
  </div>
{/each}