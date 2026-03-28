<script lang="ts">
  import Navigation from "../../components/Navigation.svelte";

  let input = $state('');
  let cursorPos = $state(0);
  let before = $derived(input.slice(0, cursorPos))
  let after = $derived(input.slice(cursorPos))
  let cursorVisible = $state(true);

  type Mode = 'login' | 'terminal' | 'password'  
  let mode = $state<Mode>('login')

  type Line = { prompt: string, value: string, br?: boolean }  
  let lines = $state<Line[]>([])

  // cursor blink :)
  let interval = setInterval(() => cursorVisible = !cursorVisible, 200)

  function resetBlink() {
    cursorVisible = true
    clearInterval(interval)
    interval = setInterval(() => cursorVisible = !cursorVisible, 200)
  }

  async function handleEnter() {
    if (mode === 'login') {
      if (input === '') {
        lines = [...lines, { prompt: 'aerytheia login:', value: '', br: false }] 
        return
      } 

      lines = [...lines, { prompt: 'aerytheia login:', value: input }]
      input = ''
      cursorPos = 0
      mode = 'password'
    } else if (mode === 'password') {
      lines = [...lines, { prompt: 'Password:', value: '', br: true }]

      console.log(input)

      input = ''
      cursorPos = 0
      
      let ok = false

      if (ok) {
        mode = 'terminal'
      } else {
        lines = [...lines, { prompt: 'Login incorrect', value: '' }]
        
        mode = 'login'
      }
    } else if (mode === 'terminal') {
      
      input = ''
      cursorPos = 0
    }
  }

  function handleKeydown(e: KeyboardEvent) {
    e.preventDefault()
    
    // skip any of these things
    if (e.altKey || e.metaKey || e.shiftKey) {
      return
    }

    // clear ctrl c on login :)
    if (e.ctrlKey && e.key.toLowerCase() === 'c') {
      if (mode === 'password' || mode === "login") {
        lines = []
        input = ''
        cursorPos = 0
        mode = "login"
      }
      return
    }

    resetBlink()

    if (e.key === 'Backspace') {
      if (cursorPos === 0) return
      input = input.slice(0, cursorPos - 1) + input.slice(cursorPos)
      cursorPos = Math.max(0, cursorPos - 1)
    } else if (e.key === 'ArrowLeft') {
      cursorPos = Math.max(0, cursorPos - 1)
    } else if (e.key === 'ArrowRight') {
      cursorPos = Math.min(input.length, cursorPos + 1)
    } else if (e.key === 'Enter') {
      handleEnter()
      // input = ''
      // cursorPos = 0
    } else if (e.key.length === 1) {
      input = input.slice(0, cursorPos) + e.key + input.slice(cursorPos)
      cursorPos++
    }
  } 


  let prompt = $derived(
    mode === 'login' ? 'aerytheia login:' :
mode === 'password' ? 'Password:' : '>'
  )

</script>


<svelte:window onkeydown={handleKeydown} />

<main class="max-w-4xl mx-auto px-4">
  <Navigation />
  {#if mode === 'login' || mode === 'password'}
    <div class="text-[var(--green)]" style="font-family: 'JetBrainsMono Nerd Font', monospace;">&lt;&lt;&lt; Welcome to aerytheia.com - tty1 &gt;&gt;&gt;</div>
    <br/>
    {#each lines as line}
      <div class="flex items-center" style="font-family: 'JetBrainsMono Nerd Font', monospace;">
        <span class="text-[var(--green)]">{line.prompt}</span>
        <span class="text-[var(--green)] ml-2">{line.value}</span>
      </div>
    {#if line.br}
      <br/>
    {/if}      
    {/each}    
  {/if}



  <!-- fake input.. -->
  <div class="flex items-center" style="font-family: 'JetBrainsMono Nerd Font', monospace;">
    <div class="text-[var(--green)] mr-2" style="font-family: 'JetBrainsMono Nerd Font', monospace;">{prompt}</div>

    {#if mode !== 'password' }
      <span class="text-[var(--green)]">{before}</span>
      <div class="bg-[var(--green)] truncate" style="width: 1ch; white-space: pre;" style:opacity={after[0] ? 1 : (cursorVisible ? 1 : 0)}>{after[0] ?? ' '}</div>
      <span class="text-[var(--green)]">{after.slice(1)}</span>
    {:else}
      <div class="bg-[var(--green)] truncate" style="width: 1ch; white-space: pre;" style:opacity={after[0] ? 1 : (cursorVisible ? 1 : 0)}>{after[0] ?? ' '}</div>
    {/if}

  </div>

</main> 