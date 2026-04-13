<script>
  import { onMount } from 'svelte'
  import { Terminal } from '@xterm/xterm';
  import { FitAddon } from '@xterm/addon-fit';
  import cowsay from 'cowsay'
  import Navigation from '../components/Navigation.svelte';
  import Footer from '../components/Footer.svelte';

  // terminal container
  let container

  // the user input
  let input = ''
  let state = 'login' // 'login', 'password', 'terminal', 'process'

  let fs = ['.???']
  
  function handleCommand(term) {
    let split = input.trim().split(' ')
    let command = split[0]
    let args = split.slice(1)


    switch (command) {
      case '':
        break
      case 'ls':
        term.writeln('')


        fs.forEach(element => {
          term.write(element + ' ')
        });

      case 'clear':
        break
      case 'hlep':
        // aperture science . com reference :D
        term.write(`\n\nIf this is an actual plea for help in response to a hazardous material spill, an explosion, a fire on your person, radiation poisoning, a choking gas of unknown origin, eye trauma resulting from the use of an emergency eye wash station on floors three, four, or eleven, an animal malfunction, or any other injurious experimental equipment failure, please remain at your workstation. A Crisis Response Team has already been mobilized to deliberate on a response to your crisis.\n\nIf you need help accessing the system, please refer to your User Handbook.`)
        break
      case 'cowsay':

      default:
        term.write(`\n${command}: command not found`)
    }

    term.write('\n\nguest> ')
  }

  function handleEnter(term) {
    // break line on enter

    if (state === 'terminal') {
      // ignore for now
      handleCommand(term)

      return
    } else if (state === 'login') {
      if (input === '') {
        term.write('\naerytheia login: ')
        return
      }

      state = 'password'
      term.write(`\nPassword: `)
    } else if (state === 'password') {

      let ok = true

      if (ok) {
        state = 'terminal'
        term.write(`\n\nguest> `)
      } else {
        term.writeln('\n\nLogin incorrect')
        term.write('aerytheia login: ')
        state = 'login'
      }
    }
  }
  
  function handleKey(term, { key, domEvent }) {
    const e = domEvent

    if (e.altKey || e.metaKey) return

    switch (e.key) {
      case 'Enter':
        handleEnter(term)
        const cmd = input.trim()
        input = ''
        
        break
      case 'Backspace':
        if (input.length > 0) {
          input = input.slice(0, -1)
          term.write('\b \b')
        }
        break
      default:
        if (e.key.length === 1) {
          input += e.key
          state === 'password' ? term.write('\u200B') : term.write(key)
        }
        break
    }
  }

  // load terminal component and hook up key event
  onMount(() => {
    // load style
    const style = getComputedStyle(document.documentElement)
    const bg = style.getPropertyValue('--bg').trim()
    const green = style.getPropertyValue('--green').trim()

    const term = new Terminal({
      fontFamily: 'JetBrainsMono Nerd Font, monospace',
      cursorBlink: true,
      convertEol: true,
      theme: {
        background: bg,
        foreground: green,
        cursor: green,
        cursorAccent: bg
      }
    })

    const fitAddon = new FitAddon()
    term.loadAddon(fitAddon)
    term.open(container)
    fitAddon.fit()
    term.focus()

    // write welcome message and initial password prompt
    term.write("<<< aerytheia - tty1 >>>\nGuest credentials:\nusername: guest\npassword: password\n\n")
    term.write("aerytheia login: ")


    term.onKey((e) => handleKey(term, e))

    return () => term.dispose()
  })
</script>

<Navigation />
<main class="max-w-6xl mx-auto px-4 pb-50">
  <div class="py-4" bind:this={container}></div>
</main> 
<Footer />