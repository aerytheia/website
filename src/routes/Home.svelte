<script>
  import Footer from "../components/Footer.svelte";
  import Navigation from "../components/Navigation.svelte";
  import Window from "../components/Window.svelte";

  let showWindow = $state(false);
</script>

<Navigation />
<main class="max-w-2xl mx-auto px-4 pb-16 text-sm" style="font-family: 'JetBrainsMono Nerd Font', monospace;">
  
  <section class="mt-5 mb-2">
    <h1 class="text-[var(--fg)] text-xl mb-1">ethan</h1>
  </section>

  <section class="text-[var(--dark-fg)] space-y-3 leading-relaxed">
    <p>hey!<img class="inline w-5" src="/yipee.gif" alt=":D" /> i'm a cs student and developer from sydney, australia.</p>
    <p>my main interests lie in systems programming, distributed systems, self-hosting, mathematics, game development and open source software.</p>
    <p>outside of programming i enjoy watching tv, sleeping and the cold. occasionally i go outside too.</p>
    <p>this website is a little spot where i can share my meaningless blog posts and experiment on fun frontend things, like this inconspicuous <button class="hover:cursor-pointer text-[var(--fg)]" onclick={() => showWindow = true}>button...</button></p>
  
    
  </section>

  <section class="mt-10">
    <p class="text-[var(--comment)] text-xs mb-2">contact</p>
    <div class="flex gap-4 text-[var(--dark-fg)]">
      <a href="https://github.com/aerytheia" class="hover:text-[var(--fg)] transition-colors">github</a>
      <a href="https://discord.com/users/540857337439846411" class="hover:text-[var(--fg)] transition-colors">discord</a>
    </div>
  </section>

  {#if showWindow}
    <Window title="Son Im Crine 😭😭😭" x={1000} y={200} minHeight={200} minWidth={200} hasSpace={false} width={300} height={300} resizeRatio={1} onClose={() => showWindow = false}>
      <img src="leorio.jpg" alt="the oreo" style="width:100%;height:100%;object-fit:contain;" />
    </Window>
  {/if}

</main>
<Footer />