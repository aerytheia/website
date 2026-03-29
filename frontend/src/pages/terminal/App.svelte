<script>
    import Navigation from "../../components/Navigation.svelte";
    import WindowManager from '../../components/WindowManager.svelte'
    import cowsay from 'cowsay'

    // window manager state
    let wm = $state()

    let commands = [
        "ls",
        "dir",
        "cowsay",
        "cowthink",
        "hlep",
        "clear",
        "rick"
    ];

    // prompt means the prefix append the output
    let loginMsg = '<<< Welcome to aerytheia.com - tty1 >>>\n'

    // username (might might this api later)
    let user = $state('guest')

    // im not making the full lfs
    // or any creation of files or folders
    let files = [];

    // login, password, terminal, etc..
    let mode = $state("login")

    // the terminal output
    let output = $state([])

    // add the original msg
    // svelte-ignore state_referenced_locally
    output.push(loginMsg)

    // captured input
    let input = $state('');
    let inputHistory = $state([])
    let historyIndex = $state(0)

    // mapping the modes to the prompts
    let prompts = $derived({
        login: "aerytheia login:",
        password: "Password:",
        terminal: `${user}@aerytheia:~$`
    })

    let prompt = $derived(prompts[mode])

    // mouse cursor
    let cursorPos = $state(0);
    let cursorVisible = $state(true); 
    let interval = setInterval(() => cursorVisible = !cursorVisible, 200)

    // the split up input
    let before = $derived(input.slice(0, cursorPos))
    let after = $derived(input.slice(cursorPos))

    function resetBlink() {
        cursorVisible = true
        clearInterval(interval)
        interval = setInterval(() => cursorVisible = !cursorVisible, 200)
    }

    function clearInput() {
        input = ''
        cursorPos = 0
    }

    // runs on clear (and in ctrl + c for non terminal modes)
    function clear() {
        output = []

        if (mode === 'password' || mode === "login") {
            output.push(loginMsg)
            mode = "login"
        }
    }

    // runs on ctrl + c
    function terminate() {
        if (mode === 'password' || mode === "login") {
            clear()
        } else {
            output.push(`${prompts.terminal} ^C\n`)
        }
    }

    async function handleCommand() {
        let split = input.trim().split(' ')
        let command = split[0]
        let args = split.slice(1)

        if (command === '') {
            output.push(`${prompts.terminal}\n`)
            return
        }

        output.push(`${prompts.terminal} ${input}`)

        if (!commands.includes(command)) {
            output.push(`${command}: command not found\n`)
            return
        }

        if (command === 'clear') {
            clear()
        } else if (command === 'hlep') {
            output.push(`If this is an actual plea for help in response to a hazardous material spill, an explosion, a fire on your person, radiation poisoning, a choking gas of unknown origin, eye trauma resulting from the use of an emergency eye wash station on floors three, four, or eleven, an animal malfunction, or any other injurious experimental equipment failure, please remain at your workstation. A Crisis Response Team has already been mobilized to deliberate on a response to your crisis.\n\nIf you need help accessing the system, please refer to your User Handbook.\n`)
        } else if (command === 'cowsay') {
            let text = []
            let opts = {}

            args.forEach((arg) => {
                if (arg.startsWith('-') && 'bdgpstwy'.includes(arg[1])) opts.mode = arg[1]
                else if (!arg.startsWith('-')) text.push(arg)
            })

            output.push(cowsay.say({ text: text.join(' '), ...opts }) + '\n')
        } else if (command = 'rick') {
            wm.openWindow(rickroll, { title: 'Cool GIF', width: 300, height: 300, resizeRatio: 1, hasSpace: false })
        }
    }

    async function handleEnter() {
        if (mode === 'terminal') {
            handleCommand()

            inputHistory.push(input)
            historyIndex = inputHistory.length
            clearInput()

            
        } else if (mode === 'login') {
            if (input === '') {
                output.push(`${prompts.login}`)
                return
            }

            output.push(`${prompts.login} ${input}`)
            clearInput()
            mode = 'password'
        } else if (mode === 'password') {
            output.push(prompts.password)

            clearInput()

            let ok = true

            if (ok) {
                mode = 'terminal'
            } else {
                output.push("Login incorrect\n")
                mode = 'login'
            }
        }
    }

    function handleKeydown(e) {
        e.preventDefault()

        // skipping these for now i guess?
        if (e.altKey || e.metaKey) {
            return
        }

        // tee hee
        resetBlink()

        // clearing

        if (e.ctrlKey && e.key.toLowerCase() === 'c') {
            terminate()
            return
        }

        if (e.key === 'Backspace') {
            // no need to go any further back twin
            if (cursorPos === 0) return
            input = input.slice(0, cursorPos - 1) + input.slice(cursorPos)
            cursorPos = Math.max(0, cursorPos - 1)
        } else if (e.key === 'ArrowLeft') {
            cursorPos = Math.max(0, cursorPos - 1)
        } else if (e.key === 'ArrowRight') {
            cursorPos = Math.min(input.length, cursorPos + 1)
        } else if (e.key === 'ArrowUp') {
            if (historyIndex <= 0) return
            historyIndex--
            input = inputHistory[historyIndex]
            cursorPos = input.length
        } else if (e.key === 'ArrowDown') {
            if (historyIndex >= inputHistory.length) return
            historyIndex++
            input = inputHistory[historyIndex] ?? ''
            cursorPos = input.length
        } else if (e.key === 'Enter') {
            handleEnter()
        } else if (e.key.length === 1) {
            input = input.slice(0, cursorPos) + e.key + input.slice(cursorPos)
            cursorPos++
        }
    }

</script>

<svelte:window onkeydown={handleKeydown} />

{#snippet rickroll()}
  <img src="/rickroll-roll.gif" alt="cool gif" style="width:100%;height:100%;object-fit:contain;" />
{/snippet}


<WindowManager bind:this={wm} />
<main class="max-w-4xl mx-auto px-4">
  <Navigation />

  {#each output as line}
    <div class="flex items-center" style="font-family: 'JetBrainsMono Nerd Font', monospace; white-space: pre">
      <pre class="text-[var(--green)]" style="font-family: 'JetBrainsMono Nerd Font', monospace; white-space: pre-wrap;">{line}</pre>
    </div>
    {#if line.endsWith('\n')}<br/>{/if}
  {/each}

  <!-- fake input.. -->
  <div class="flex items-center" style="font-family: 'JetBrainsMono Nerd Font', monospace; white-space: pre">
    <pre class="text-[var(--green)]" style="font-family: 'JetBrainsMono Nerd Font', monospace; ">{prompt} </pre>

    {#if mode !== 'password' }
      <span class="text-[var(--green)]">{before}</span>
      <div class="bg-[var(--green)] truncate" style="width: 1ch; white-space: pre;" style:opacity={after[0] ? 1 : (cursorVisible ? 1 : 0)}>{after[0] ?? ' '}</div>
      <span class="text-[var(--green)]">{after.slice(1)}</span>
    {:else}
      <div class="bg-[var(--green)] truncate" style="width: 1ch; white-space: pre;" style:opacity={after[0] ? 1 : (cursorVisible ? 1 : 0)}>{after[0] ?? ' '}</div>
    {/if}
  </div>

</main> 