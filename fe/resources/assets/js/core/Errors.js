
class Errors {

	constructor() {
		this.errors = {};
	}

	get(field) {
		if (this.errors[field]) {
			return this.errors[field];
		}
	}

	record(errors) {
		return this.errors = errors;
	}

	any() {
		return Object.keys(this.errors).length > 0;
	}

	clear(field) {
		if (field) {
			delete this.errors[field];
			return;
		}
		this.errors = {};
		
	}

	has(field){
		if (field in this.errors) {
			return true;
		} else {
			return false;
		}
	}
}

export default Errors;
