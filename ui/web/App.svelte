<script>
	import { getCookie } from './lib/cookie.js';
	import { callCheckConnexion } from "./lib/network.js";

	import Login from './functions/Login.svelte';
	import Dashboard from './functions/Dashboard.svelte';

	// pas de jeton => l'utilisateur n'est pas connecté
	// si on a un jeton, on vérifie auprès du serveur
	let connectedStatus = true;
	if(getCookie("jeton") === "") connectedStatus = false; 
	else callCheckConnexion().then(response => { connectedStatus = response; });	
	// vérifie périodiquement que la connexion est toujours ok
	function periodicCheck() {
		callCheckConnexion().then(response => { 
			connectedStatus = response; 
			if (connectedStatus) {
				setTimeout(periodicCheck, 5000);
			}
		});	
	}
	$: if(connectedStatus) periodicCheck();
</script>

{#if !connectedStatus}
	<Login bind:connectedStatus/>
{:else}
	<Dashboard/>
{/if}
