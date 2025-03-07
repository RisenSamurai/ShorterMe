import type { PageServerLoad } from './$types';
import { redirect, error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, fetch }) => {
    const { id } = params;

    const response = await fetch(`http://backend:8080/redirect/${id}`);

    if (response.ok) {
        const data = await response.json();
        if (data.url) {
            throw redirect(302, data.url); // Correct frontend redirect
        } else {
            throw error(404, 'Short URL not found');
        }
    } else {
        throw error(response.status, 'Failed to fetch URL');
    }
};
