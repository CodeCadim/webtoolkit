<script>
  import Error from "../components/Error.svelte";
  import Input from "../components/Input.svelte";
  import Button from "../components/Button.svelte";

  export let connectedStatus = false;

  let errMessage;
  let credential = {username: "", password: ""};

  async function login() {
    errMessage = "";
	try {
		let response = await fetch("/login", {
		  method: "POST",
		  body: JSON.stringify(credential),
		});
		if (response.ok) {
		  connectedStatus = true;
		} else {
		  connectedStatus = false;
		  errMessage = "Echec de la connexion. Veuillez vérifier votre identifiant et votre mot de passe.";
		}
	}
	catch(error) {
		errMessage = "Problème réseau: impossible de contacter le serveur. ";
	}
  }
</script>

<section >
  <main class="login">

{#if errMessage}
<Error>{errMessage}</Error>
{/if}

    <h1>
      <span class="icon">
        <i class="las la-cube la-fw" />
      </span>
      <span>Webtoolkit</span>
    </h1>
    <form on:submit|preventDefault={login} action="/login" method="post" class="box">

	  <Input libelle="Identifiant" name="username" icon="las la-user" bind:value={credential.username} autofocus="autofocus" />

	  <Input type="password" libelle="Mot de passe" name="password" icon="las la-unlock" bind:value="{credential.password}" />

      <div class="field">
        <Button>
          Connecter
        </Button>
      </div>
    </form>
  </main>
</section>
