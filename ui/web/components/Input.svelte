<script>
	export let value;
	export let libelle;
	export let type="text";
	export let name;
	export let icon;
	export let autofocus = "";
	export let placeholder = "";

	// hack to pass a type input (otherwise, svelte complain about type attribute cannot be dynamic bla bla)
	function typeAction(node) {
		node.type = type;
	}
</script>

<label for="{name}" class="label">{libelle}</label>
<div class="field icon">
  <!-- svelte-ignore a11y-autofocus-->
  <input use:typeAction {name} {placeholder} {autofocus} bind:value />
	<i class="{icon}" />
</div>
