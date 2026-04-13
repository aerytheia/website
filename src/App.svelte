<script>
  import Router from 'svelte-spa-router'

  import Home from './routes/Home.svelte'
  import Terminal from './routes/Terminal.svelte';

  const routes = {
      '/': Home,
      '/terminal': Terminal,
  }  
</script>

<Router {routes} />