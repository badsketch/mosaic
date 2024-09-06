import type { Actions } from './$types';

export const actions = {
	default: async ({ request }) => {
		const formData = Object.fromEntries(await request.formData());
		const res = await fetch('http://localhost:5000/', { method: 'POST', body: formData });
		const json = await res.json();
		console.log('response', json);
		return { data: json };
	}
} satisfies Actions;
