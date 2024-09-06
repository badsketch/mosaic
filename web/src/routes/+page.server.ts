import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const res = await fetch('http://localhost:5000/');
	const json = await res.json();
	return { data: json };
};
