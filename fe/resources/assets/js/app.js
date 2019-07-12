
import Vue from "vue";
import axios from "axios";
import Form from "./core/Form"
import Example from "./components/Example";

window.axios = axios;
window.Form = Form;

new Vue({

	el: "#app",

	components: {
		Example,
	},

	data: {
		form: new Form({
			name: "",
			description: "",
		}),
	},

	methods: {
		onSubmit() {
			this.form.post("axios");
		},

	},

	mounted() {
		axios.get("/mounted").then(response => console.log("from /mounted:", response.data));
	},

});
