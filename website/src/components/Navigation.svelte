<script>
  import {router} from 'svelte-spa-router'

  const links = ['home', 'blog', 'guestbook', 'projects', 'terminal']
</script>


<nav class="w-full justify-center py-1 flex z-1 gap-4" style="font-family: 'JetBrainsMono Nerd Font', monospace;">
  {#each links as link, i}
    {@const href = link === 'home' ? '/' : `/${link}`}
    <a href={`#${href}`} 
       class="transition-colors {router.location === href ? 'text-[var(--fg)]' : 'text-[var(--comment)] hover:text-[var(--fg)]'}">
      {link}
    </a>
  {/each}
</nav>
<hr class="border-t border-[var(--comment)]">