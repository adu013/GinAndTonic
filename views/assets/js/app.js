document.addEventListener('DOMContentLoaded', () => {
    const copyBtn = document.getElementById('copy-btn');
    const copyText = document.getElementById('copy-text');

    if (copyBtn && copyText) {
        copyBtn.addEventListener('click', async () => {
            const textToCopy = copyText.innerText;

            try {
                // Copy package text to system clipboard
                await navigator.clipboard.writeText(textToCopy);

                // Visual feedback update
                const originalText = copyText.innerText;
                copyText.innerText = "Copied! 🎉";
                copyBtn.style.borderColor = "var(--accent)";

                // Revert button styling back to default after 2 seconds
                setTimeout(() => {
                    copyText.innerText = originalText;
                    copyBtn.style.borderColor = "var(--border)";
                }, 2000);

            } catch (err) {
                console.error('Failed to copy text: ', err);
            }
        });
    }
});
