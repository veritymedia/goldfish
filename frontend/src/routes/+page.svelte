<h1 class="text-[red]">Welcome to SvelteKit</h1>
<p>{fromBackend}</p>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>


<script>
import {InitClipboard} from "../lib/wailsjs/go/main/App.js"

let fromBackend = "Loading..."

function loadFromBackend() {
    InitClipboard().then((result) => {
        fromBackend = result
    })
}

loadFromBackend()

</script>
