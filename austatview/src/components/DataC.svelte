<script>
	import CollapsibleSection from './CollapsibleSection.svelte'
  export let name;
  export let description;
</script>
<li>
<CollapsibleSection headerText={name} >
	<div class="content">
  {description}
	</div>
</CollapsibleSection>
</li>
<style>
	section {
		width: 300px;
	}

	.content {
		background-color: #f4f4f4;
		padding: 0.5em;
	}

</style>
