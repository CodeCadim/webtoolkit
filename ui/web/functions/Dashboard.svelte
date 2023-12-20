<script>
  import Bonjour from "./Bonjour.svelte";
  import List from "./List.svelte";

  let Menu = [
    { id: "bonjour", component: Bonjour, icon: "la-handshake", label: "Bonjour" },
	{ id: "list", component: List, icon: "la-users", label: "Utilisateurs" }
  ];

  let showNavbarMenu = false;
  let actualMenu = Menu[0];

  function activate(menuID) {
    actualMenu = Menu.find(elem => elem.id === menuID);
    showNavbarMenu = !showNavbarMenu;
  }
</script>

<header>
  <div class="item">
    <label for="toggler" id="togglerlabel">
      <span class="icon">
        <i class="las la-bars" />
      </span>
    </label>
    <div class="brand">
      <span class="icon">
        <i class="las la-cube la-fw" />
      </span>
      <span>Webtoolkit</span>
    </div>
  </div>
  <div class="end item">
    <div class="profile">
      <label for="profilebox">
        <span class="icon">
          <i class="las la-user" />
        </span>
      </label>
      <input type="checkbox" id="profilebox" checked="checked" />
      <div class="profilemenu">
        <a href="#profile">Modifier son profile</a>
        <a href="/logout">Se déconnecter</a>
      </div>
    </div>
  </div>
</header>
<section>
  <aside>
    <input type="checkbox" id="toggler" />
    <nav>
      {#each Menu as item (item.id)}
        <a href="#{item.id}" class="item" class:selected={actualMenu.id==item.id} on:click={() => activate(item.id)}>
          <span class="icon">
            <i class="las la-fw {item.icon}" />
          </span>
          <span>{item.label}</span>
        </a>
      {/each}
    </nav>
  </aside>
  <main>
    <svelte:component this={actualMenu.component} />
  </main>
</section>

<style>
  @media screen and (max-width: 1023px) {
    .navbar-menu {
      position: absolute;
      max-width: max-content;
      right: 0;
    }
  }
</style>
