<script>
  const links = ['home', 'blog', 'guestbook', 'projects', 'terminal', 'about']
</script>

<nav class="w-full flex justify-center gap-4 py-1 text-sm">
    {#each links as link}
        <a href={link === 'home' ? '/' : `/${link}`} class="hover:text-[var(--fg)] text-[var(--comment)] transition-colors" style="font-family: 'JetBrainsMono Nerd Font', monospace;">
            {link}
        </a>
    {/each}
</nav>