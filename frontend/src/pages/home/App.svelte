<script>
  import Navigation from "../../components/Navigation.svelte";
  import Window from "../../components/Window.svelte";

  let showWindow = $state(true);
</script>

<main class="max-w-4xl mx-auto px-4">
  <Navigation />
  {#if showWindow}
    <Window title="Cool GIF" resizeRatio={1} width={300} height={300} hasSpace={false}>
      <img src="/rickroll-roll.gif" alt="cool gif" style="width: 100%; height: 100%; object-fit: contain;" />
    </Window>
  {/if}
</main>