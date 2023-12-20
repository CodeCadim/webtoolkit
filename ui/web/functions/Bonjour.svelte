<script>
  import * as net from "../lib/network.js";
  import Error from "../components/Error.svelte";
  import Input from "../components/Input.svelte";
  import Button from "../components/Button.svelte";

  let resultat;
  let nom = "";
  let errMessage = "";

  async function bonjour() {
	errMessage = "";
	if(nom == "") {
		errMessage = "Un prénom doit être renseigné";
		return;
	}
    let res = await net.callBonjour(nom);
	if (res.error != null) {
		errMessage = res.error;
		return
	}
	resultat = res.data;
  }
</script>

<h1>Bonjour</h1>

{#if errMessage}
<Error>{errMessage}</Error>
{/if}

<form on:submit|preventDefault={bonjour}>
  <Input libelle="Entrer un prénom" placeholder="Prénom..." bind:value={nom} name="name" icon="las la-user" autofocus="autofocus" />
  <Button>
	Bonjour
  </Button>
</form>
{#if resultat}
<hr />
{resultat}
{/if}
