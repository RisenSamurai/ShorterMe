export const actions = {
    default: async ({ request }) => {
        const formData = await request.formData();
        const url = formData.get('url');



        try {
            const response = await fetch("http://backend:8080/shorten", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ url }),
            });

            const data = await response.json();

            return {
                success: response.ok,
                shortUrl: data.url || null,
                message: data.message || "Unable to generate short URL.",
            };
        } catch (error) {
            console.error("Fetch Error:", error);
            return {
                success: false,
                message: "A server error occurred. Please try again later.",
            };
        }
    },
};
