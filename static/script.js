document.addEventListener("DOMContentLoaded", function () {
    const toggleBtn = document.getElementById("darkToggle");

    // Apply saved theme on page load
    const isDark = localStorage.getItem("darkMode") === "true";
    if (isDark) {
        enableDarkMode();
    }

    toggleBtn.addEventListener("click", () => {
        const currentlyDark = document.body.classList.contains("dark-mode");
        if (currentlyDark) {
            disableDarkMode();
            localStorage.setItem("darkMode", "false");
        } else {
            enableDarkMode();
            localStorage.setItem("darkMode", "true");
        }
    });

    function enableDarkMode() {
        document.body.classList.add("dark-mode");

        document.querySelectorAll(".card").forEach(card =>
            card.classList.add("dark-mode")
        );

        document.querySelectorAll(".bg-light").forEach(bg =>
            bg.classList.add("dark-mode")
        );

        toggleBtn.classList.add("dark-mode");
    }

    function disableDarkMode() {
        document.body.classList.remove("dark-mode");

        document.querySelectorAll(".card").forEach(card =>
            card.classList.remove("dark-mode")
        );

        document.querySelectorAll(".bg-light").forEach(bg =>
            bg.classList.remove("dark-mode")
        );

        toggleBtn.classList.remove("dark-mode");
    }
});
