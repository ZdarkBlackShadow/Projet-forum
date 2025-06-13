document.addEventListener('DOMContentLoaded', () => {
    const categoryLinks = document.querySelectorAll('.categories-nav li');
    const serverItems = document.querySelectorAll('.server-item');
    const searchBar = document.querySelector('.search-bar');

    // Category filtering
    categoryLinks.forEach(link => {
        link.addEventListener('click', function() {
            const filter = this.getAttribute('data-filter');

            categoryLinks.forEach(l => l.classList.remove('active-category'));
            this.classList.add('active-category');

            serverItems.forEach(item => {
                if (filter === 'all' || item.getAttribute('data-category') === filter) {
                    item.style.display = '';
                } else {
                    item.style.display = 'none';
                }
            });
        });
    });

    // Search bar functionality
    if (searchBar) {
        searchBar.addEventListener('keyup', function() {
            const searchTerm = this.value.toLowerCase();
            const activeFilter = document.querySelector('.categories-nav li.active-category')?.getAttribute('data-filter') || 'all';


            serverItems.forEach(item => {
                const itemName = item.textContent.toLowerCase();
                const itemCategory = item.getAttribute('data-category');

                const matchesSearch = itemName.includes(searchTerm);
                const matchesCategory = (activeFilter === 'all' || itemCategory === activeFilter);

                if (matchesSearch && matchesCategory) {
                    item.style.display = '';
                } else {
                    item.style.display = 'none';
                }
            });
        });
    }
});