// IDGEN Frontend Application

(function() {
    'use strict';

    // State
    let currentData = null;
    let batchData = [];

    // DOM Elements
    const themeToggle = document.getElementById('theme-toggle');
    const generateBtn = document.getElementById('generate-btn');
    const copyAllBtn = document.getElementById('copy-all-btn');
    const batchCountSelect = document.getElementById('batch-count');
    const exportCsvBtn = document.getElementById('export-csv-btn');
    const singleView = document.getElementById('single-view');
    const batchView = document.getElementById('batch-view');
    const batchTable = document.getElementById('batch-table').querySelector('tbody');
    const toast = document.getElementById('toast');

    // Initialize
    function init() {
        // Load saved theme
        const savedTheme = localStorage.getItem('idgen-theme');
        if (savedTheme) {
            document.documentElement.setAttribute('data-theme', savedTheme);
        }

        // Bind events
        themeToggle.addEventListener('click', toggleTheme);
        generateBtn.addEventListener('click', generate);
        copyAllBtn.addEventListener('click', copyAll);
        exportCsvBtn.addEventListener('click', exportCSV);
        batchCountSelect.addEventListener('change', onBatchCountChange);

        // Bind copy buttons
        document.querySelectorAll('.copy-btn[data-target]').forEach(btn => {
            btn.addEventListener('click', () => copyField(btn.dataset.target));
        });

        // Initial generate
        generate();
    }

    // Theme toggle
    function toggleTheme() {
        const html = document.documentElement;
        const currentTheme = html.getAttribute('data-theme');
        const newTheme = currentTheme === 'cyber' ? 'terminal' : 'cyber';
        html.setAttribute('data-theme', newTheme);
        localStorage.setItem('idgen-theme', newTheme);
    }

    // Generate data
    async function generate() {
        const count = parseInt(batchCountSelect.value, 10);

        try {
            generateBtn.disabled = true;
            generateBtn.querySelector('span:last-child').textContent = 'Loading...';

            if (count === 1) {
                const response = await fetch('/api/v1/generate');
                if (!response.ok) throw new Error('API error');
                currentData = await response.json();
                batchData = [currentData];
                updateSingleView(currentData);
                showSingleView();
            } else {
                const response = await fetch(`/api/v1/batch?count=${count}`);
                if (!response.ok) throw new Error('API error');
                const result = await response.json();
                batchData = result.data;
                currentData = batchData[0];
                updateSingleView(currentData);
                updateBatchTable(batchData);
                showBatchView();
            }
        } catch (error) {
            console.error('Generate failed:', error);
            showToast('Generation failed');
        } finally {
            generateBtn.disabled = false;
            generateBtn.querySelector('span:last-child').textContent = 'Generate';
        }
    }

    // Update single view
    function updateSingleView(data) {
        document.querySelector('[data-field="name"]').textContent = data.name;
        document.querySelector('[data-field="idno"]').textContent = data.idno;
        document.querySelector('[data-field="mobile"]').textContent = data.mobile;
        document.querySelector('[data-field="bank"]').textContent = data.bank;
        document.querySelector('[data-field="email"]').textContent = data.email;
        document.querySelector('[data-field="address"]').textContent = data.address;
    }

    // Update batch table
    function updateBatchTable(data) {
        batchTable.innerHTML = '';
        const fields = ['name', 'idno', 'mobile', 'bank', 'email', 'address'];

        data.forEach((item, index) => {
            const row = document.createElement('tr');

            // Index column
            const indexCell = document.createElement('td');
            indexCell.textContent = index + 1;
            row.appendChild(indexCell);

            // Data columns (click to copy)
            fields.forEach(field => {
                const cell = document.createElement('td');
                cell.className = 'copyable';
                cell.setAttribute('title', 'Click to copy: ' + item[field]);
                cell.setAttribute('data-value', item[field]);
                cell.textContent = item[field];
                row.appendChild(cell);
            });

            // Row copy button column
            const actionCell = document.createElement('td');
            actionCell.innerHTML = `
                <button class="copy-btn" title="Copy Row">
                    <span class="copy-icon">&#x2398;</span>
                </button>
            `;
            row.appendChild(actionCell);

            batchTable.appendChild(row);
        });

        // Bind cell copy events
        bindCellCopyEvents();
    }

    // Bind cell copy events
    function bindCellCopyEvents() {
        document.querySelectorAll('td.copyable').forEach(cell => {
            // Click on cell or button to copy
            cell.addEventListener('click', (e) => {
                const value = cell.getAttribute('data-value');
                copyToClipboard(value);

                // Visual feedback
                cell.classList.add('copied');
                setTimeout(() => cell.classList.remove('copied'), 500);

                showToast('Copied');
            });
        });

        // Bind row copy buttons
        document.querySelectorAll('#batch-table tbody tr').forEach((row, index) => {
            const rowBtn = row.querySelector('td:last-child .copy-btn');
            if (rowBtn) {
                rowBtn.addEventListener('click', (e) => {
                    e.stopPropagation();
                    copyRowByIndex(index);
                });
            }
        });
    }

    // Copy row by index
    function copyRowByIndex(index) {
        if (!batchData[index]) return;

        const item = batchData[index];
        const text = [
            item.name,
            item.idno,
            item.mobile,
            item.bank,
            item.email,
            item.address
        ].join('\n');

        copyToClipboard(text);
        showToast('Row copied');
    }

    // Show/hide views
    function showSingleView() {
        batchView.style.display = 'none';
    }

    function showBatchView() {
        batchView.style.display = 'block';
    }

    // Batch count change handler
    function onBatchCountChange() {
        const count = parseInt(batchCountSelect.value, 10);
        if (count === 1) {
            showSingleView();
        }
    }

    // Copy single field
    function copyField(field) {
        if (!currentData) return;

        const value = currentData[field];
        copyToClipboard(value);

        // Visual feedback
        const btn = document.querySelector(`.copy-btn[data-target="${field}"]`);
        btn.classList.add('copied');
        setTimeout(() => btn.classList.remove('copied'), 500);

        showToast('Copied');
    }

    // Copy all fields
    function copyAll() {
        if (!currentData) return;

        const text = [
            currentData.name,
            currentData.idno,
            currentData.mobile,
            currentData.bank,
            currentData.email,
            currentData.address
        ].join('\n');

        copyToClipboard(text);
        showToast('All copied');
    }

    // Export CSV
    function exportCSV() {
        const count = parseInt(batchCountSelect.value, 10);
        window.location.href = `/api/v1/export?count=${count}`;
    }

    // Copy to clipboard utility
    function copyToClipboard(text) {
        if (navigator.clipboard && navigator.clipboard.writeText) {
            navigator.clipboard.writeText(text).catch(err => {
                console.error('Clipboard write failed:', err);
                fallbackCopy(text);
            });
        } else {
            fallbackCopy(text);
        }
    }

    // Fallback copy for older browsers
    function fallbackCopy(text) {
        const textarea = document.createElement('textarea');
        textarea.value = text;
        textarea.style.position = 'fixed';
        textarea.style.opacity = '0';
        document.body.appendChild(textarea);
        textarea.select();
        try {
            document.execCommand('copy');
        } catch (err) {
            console.error('Fallback copy failed:', err);
        }
        document.body.removeChild(textarea);
    }

    // Toast notification
    function showToast(message) {
        toast.textContent = message;
        toast.classList.add('show');
        setTimeout(() => toast.classList.remove('show'), 1500);
    }

    // Start app
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }
})();
