<script lang="ts">
    export let form;

    // Hydration-safe reactive variables
    let shortUrl = null;
    let errorMessage = null;

    // Handle response from the server
    if (form) {
        if (form.success) {
            shortUrl = form.shortUrl;
        } else {
            errorMessage = form.message || "An unexpected error occurred.";
        }
    }
</script>

<section class="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-900">
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8 w-full max-w-md">
        <div class="text-center mb-6">
            <h1 class="text-3xl font-bold uppercase text-rose-500">Shorter.me</h1>
        </div>

        <form method="POST" class="space-y-4">
            <div>
                <label for="url" class="block text-sm font-bold text-gray-700 dark:text-gray-300">Paste your URL</label>
                <input
                        id="url"
                        type="text"
                        name="url"
                        placeholder="Enter your URL..."
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-rose-500 focus:border-rose-500"
                />
            </div>


            <button
                    type="submit"
                    class="w-full py-2 px-4 bg-rose-500 hover:bg-rose-600 text-white font-bold rounded-md focus:outline-none focus:ring-2 focus:ring-rose-500"
            >
                Submit
            </button>

            {#if form}
                {#if form.success}
                    <p class="text-center text-green-500">
                        Your shortened URL: <a href={form.shortUrl} target="_blank" class="underline hover:text-green-600">{form.shortUrl} </a>
                    </p>
                {:else}
                    <p class="text-center text-red-500">Error: {form.message}</p>
                {/if}
            {/if}
        </form>
    </div>
</section>
