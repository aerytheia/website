<script>
  const links = ['home', 'blog', 'guestbook', 'projects', 'terminal', 'about']
</script>

<nav class="w-full flex justify-center gap-4 py-1 text-sm">
    {#each links as link}
        <a href="/{link}" class="hover:text-[var(--comment)] text-[var(--fg)] transition-colors">
        {link}
        </a>
    {/each}
</nav>