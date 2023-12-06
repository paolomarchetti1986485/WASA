<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			user: null,
		}
	},
	methods: {
		async delete(id){
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/user" + id);
				await this.refresh();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/user");
				this.user = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ul>
			<li v-for="u in user">
				ID: {{u.userId}}, 
				Username: {{u.username}}
				<a href="javascript:" @click="removeUserHandler(u.userId)">[DELETE]</a>
			</li>
		</ul>
	</div>
</template>

<style>
</style>
