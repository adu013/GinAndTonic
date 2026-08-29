document.addEventListener('DOMContentLoaded', () => {
    const switcher = document.getElementById('theme-switcher');
    const themeIcon = document.getElementById('theme-icon');
    const themeText = document.getElementById('theme-text');

    if (switcher) {
        // Check if user already has a saved preference, otherwise default to dark mode
        const currentTheme = localStorage.getItem('theme') || 'dark';
        document.documentElement.setAttribute('data-theme', currentTheme);
        updateToggleUI(currentTheme);

        switcher.addEventListener('click', () => {
            let targetTheme = 'dark';
            if (document.documentElement.getAttribute('data-theme') === 'dark') {
                targetTheme = 'light';
            }

            document.documentElement.setAttribute('data-theme', targetTheme);
            localStorage.setItem('theme', targetTheme);
            updateToggleUI(targetTheme);
        });
    }

    function updateToggleUI(theme) {
        if (!themeIcon || !themeText) return;

        if (theme === 'light') {
            themeIcon.innerText = '☀️';
            themeText.innerText = 'Light Mode';
        } else {
            themeIcon.innerText = '🌙';
            themeText.innerText = 'Dark Mode';
        }
    }
});
