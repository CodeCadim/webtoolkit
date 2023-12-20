<script>
	import * as net from "../lib/network.js";
    import Error from "../components/Error.svelte";

	let users = []; 
	let errMessage;
	async function list() {
		let res = await net.callGetUsers()
		if (res.error != null) {
			errMessage = res.error;
			return
		}
		users = res.data;
	}
	list();
</script>

{#if errMessage}
<Error>{errMessage}</Error>
{/if}

<h1>Liste des utilisateurs</h1>

<div class="grid grid-cols-3 gap-4">
{#each users as user}
	<div> {user.Username} </div>
	<div> {user.Role} </div>
	<div> {user.Realm} </div>
{/each}
</div>
