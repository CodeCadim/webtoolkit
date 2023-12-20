export async function callApi(url) {
    try {
        let response = await fetch(url);
		let data = await response.json()
        if(response.ok) {
            return {
				error: null,
				data: data
			};
        }
		return {
				error: response.status,
				data: null
		};
    } catch (error) {
		return {
				error: error,
				data: null
		};
    }
}

export async function callCheckConnexion() {
	let response;
	try {
		response = await fetch("check");
		return response.ok;
	} catch (error) {
        console.log("Erreur réseau " + error);
		return null;
	}
}

export async function callBonjour(nom) {
    return await callApi("/hello?nom="+nom);
}

export async function callGetUsers() {
	return await callApi("/list");
}
