// IDGEN Frontend Application

(function() {
    'use strict';

    // ==================== Canvas Background System ====================
    const CanvasBackground = {
        canvas: null,
        ctx: null,
        particles: [],
        matrixColumns: [],
        mouse: { x: null, y: null },
        animationId: null,
        currentTheme: null,

        // Cyberpunk theme colors
        cyberColors: [
            { r: 0, g: 245, b: 255 },   // Cyan
            { r: 255, g: 0, b: 255 },   // Magenta
            { r: 0, g: 200, b: 255 },   // Light blue
            { r: 200, g: 0, b: 255 },   // Purple
        ],

        // Matrix characters
        matrixChars: '01アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン',

        init() {
            this.canvas = document.getElementById('bg-canvas');
            if (!this.canvas) return;

            this.ctx = this.canvas.getContext('2d');
            this.resize();

            // Event listeners
            window.addEventListener('resize', () => this.resize());
            window.addEventListener('mousemove', (e) => {
                this.mouse.x = e.clientX;
                this.mouse.y = e.clientY;
            });
            window.addEventListener('mouseout', () => {
                this.mouse.x = null;
                this.mouse.y = null;
            });

            // Start animation
            this.currentTheme = document.documentElement.getAttribute('data-theme');
            this.setupTheme();
            this.animate();
        },

        resize() {
            this.canvas.width = window.innerWidth;
            this.canvas.height = window.innerHeight;
            this.setupTheme();
        },

        setupTheme() {
            const theme = document.documentElement.getAttribute('data-theme');
            if (theme === 'cyber') {
                this.setupCyberParticles();
            } else {
                this.setupMatrixRain();
            }
        },

        // ==================== Cyberpunk Particle System ====================
        setupCyberParticles() {
            this.particles = [];
            const particleCount = Math.min(120, Math.floor((this.canvas.width * this.canvas.height) / 8000));

            for (let i = 0; i < particleCount; i++) {
                this.particles.push({
                    x: Math.random() * this.canvas.width,
                    y: Math.random() * this.canvas.height,
                    vx: (Math.random() - 0.5) * 0.3,
                    vy: (Math.random() - 0.5) * 0.3,
                    baseSpeed: Math.random() * 0.08 + 0.04,
                    angle: Math.random() * Math.PI * 2,
                    angleSpeed: (Math.random() - 0.5) * 0.005,
                    size: Math.random() * 2 + 1,
                    color: this.cyberColors[Math.floor(Math.random() * this.cyberColors.length)],
                    alpha: Math.random() * 0.5 + 0.3,
                    pulse: Math.random() * Math.PI * 2,
                });
            }
        },

        drawCyberBackground() {
            const ctx = this.ctx;
            const w = this.canvas.width;
            const h = this.canvas.height;

            // Clear with gradient background
            const gradient = ctx.createLinearGradient(0, 0, w, h);
            gradient.addColorStop(0, '#0a0a0f');
            gradient.addColorStop(0.5, '#12121f');
            gradient.addColorStop(1, '#0a0a0f');
            ctx.fillStyle = gradient;
            ctx.fillRect(0, 0, w, h);

            // Draw grid
            ctx.strokeStyle = 'rgba(0, 245, 255, 0.03)';
            ctx.lineWidth = 1;
            const gridSize = 60;

            for (let x = 0; x <= w; x += gridSize) {
                ctx.beginPath();
                ctx.moveTo(x, 0);
                ctx.lineTo(x, h);
                ctx.stroke();
            }
            for (let y = 0; y <= h; y += gridSize) {
                ctx.beginPath();
                ctx.moveTo(0, y);
                ctx.lineTo(w, y);
                ctx.stroke();
            }

            // Update and draw particles
            const connectionDistance = 150;
            const mouseRadius = 120;

            for (let i = 0; i < this.particles.length; i++) {
                const p = this.particles[i];

                // Wandering movement - slow drift
                p.angle += p.angleSpeed;
                p.vx += Math.cos(p.angle) * p.baseSpeed * 0.05;
                p.vy += Math.sin(p.angle) * p.baseSpeed * 0.05;

                // Mouse interaction
                if (this.mouse.x !== null && this.mouse.y !== null) {
                    const dx = p.x - this.mouse.x;
                    const dy = p.y - this.mouse.y;
                    const dist = Math.sqrt(dx * dx + dy * dy);
                    if (dist < mouseRadius) {
                        const force = (mouseRadius - dist) / mouseRadius;
                        p.vx += (dx / dist) * force * 0.8;
                        p.vy += (dy / dist) * force * 0.8;
                    }
                }

                // Update position
                p.x += p.vx;
                p.y += p.vy;

                // Damping - keeps speed under control
                p.vx *= 0.96;
                p.vy *= 0.96;

                // Bounce off edges
                if (p.x < 0 || p.x > w) {
                    p.vx *= -1;
                    p.angle = Math.PI - p.angle;
                }
                if (p.y < 0 || p.y > h) {
                    p.vy *= -1;
                    p.angle = -p.angle;
                }
                p.x = Math.max(0, Math.min(w, p.x));
                p.y = Math.max(0, Math.min(h, p.y));

                // Pulse effect
                p.pulse += 0.02;
                const pulseAlpha = p.alpha + Math.sin(p.pulse) * 0.15;

                // Draw connections
                for (let j = i + 1; j < this.particles.length; j++) {
                    const p2 = this.particles[j];
                    const dx = p.x - p2.x;
                    const dy = p.y - p2.y;
                    const dist = Math.sqrt(dx * dx + dy * dy);

                    if (dist < connectionDistance) {
                        const alpha = (1 - dist / connectionDistance) * 0.3;
                        ctx.strokeStyle = `rgba(${p.color.r}, ${p.color.g}, ${p.color.b}, ${alpha})`;
                        ctx.lineWidth = 0.5;
                        ctx.beginPath();
                        ctx.moveTo(p.x, p.y);
                        ctx.lineTo(p2.x, p2.y);
                        ctx.stroke();
                    }
                }

                // Draw particle with glow
                const glowSize = p.size * 4;
                const glow = ctx.createRadialGradient(p.x, p.y, 0, p.x, p.y, glowSize);
                glow.addColorStop(0, `rgba(${p.color.r}, ${p.color.g}, ${p.color.b}, ${pulseAlpha})`);
                glow.addColorStop(0.4, `rgba(${p.color.r}, ${p.color.g}, ${p.color.b}, ${pulseAlpha * 0.3})`);
                glow.addColorStop(1, 'transparent');

                ctx.fillStyle = glow;
                ctx.beginPath();
                ctx.arc(p.x, p.y, glowSize, 0, Math.PI * 2);
                ctx.fill();

                // Draw core
                ctx.fillStyle = `rgba(${p.color.r}, ${p.color.g}, ${p.color.b}, ${pulseAlpha + 0.3})`;
                ctx.beginPath();
                ctx.arc(p.x, p.y, p.size, 0, Math.PI * 2);
                ctx.fill();
            }

            // Draw mouse glow
            if (this.mouse.x !== null && this.mouse.y !== null) {
                const mouseGlow = ctx.createRadialGradient(
                    this.mouse.x, this.mouse.y, 0,
                    this.mouse.x, this.mouse.y, mouseRadius
                );
                mouseGlow.addColorStop(0, 'rgba(0, 245, 255, 0.1)');
                mouseGlow.addColorStop(0.5, 'rgba(255, 0, 255, 0.05)');
                mouseGlow.addColorStop(1, 'transparent');
                ctx.fillStyle = mouseGlow;
                ctx.beginPath();
                ctx.arc(this.mouse.x, this.mouse.y, mouseRadius, 0, Math.PI * 2);
                ctx.fill();
            }
        },

        // ==================== Matrix Rain Effect ====================
        setupMatrixRain() {
            this.matrixColumns = [];
            const fontSize = 14;
            const columns = Math.floor(this.canvas.width / fontSize);

            for (let i = 0; i < columns; i++) {
                this.matrixColumns.push({
                    x: i * fontSize,
                    y: Math.random() * this.canvas.height,
                    speed: Math.random() * 3 + 2,
                    fontSize: fontSize,
                    chars: [],
                    length: Math.floor(Math.random() * 15) + 5,
                });

                // Initialize characters for this column
                for (let j = 0; j < this.matrixColumns[i].length; j++) {
                    this.matrixColumns[i].chars.push({
                        char: this.matrixChars[Math.floor(Math.random() * this.matrixChars.length)],
                        changeTimer: Math.random() * 10,
                    });
                }
            }
        },

        drawMatrixBackground() {
            const ctx = this.ctx;
            const w = this.canvas.width;
            const h = this.canvas.height;

            // Semi-transparent black for trail effect
            ctx.fillStyle = 'rgba(0, 0, 0, 0.1)';
            ctx.fillRect(0, 0, w, h);

            ctx.font = '14px "JetBrains Mono", monospace';

            for (const col of this.matrixColumns) {
                // Update position
                col.y += col.speed;

                // Reset if off screen
                if (col.y - col.length * col.fontSize > h) {
                    col.y = -col.length * col.fontSize;
                    col.speed = Math.random() * 3 + 2;
                }

                // Draw characters
                for (let i = 0; i < col.chars.length; i++) {
                    const charData = col.chars[i];
                    const y = col.y - i * col.fontSize;

                    if (y < -col.fontSize || y > h + col.fontSize) continue;

                    // Randomly change character
                    charData.changeTimer -= 0.1;
                    if (charData.changeTimer <= 0) {
                        charData.char = this.matrixChars[Math.floor(Math.random() * this.matrixChars.length)];
                        charData.changeTimer = Math.random() * 20 + 5;
                    }

                    // Calculate alpha based on position in stream
                    let alpha;
                    if (i === 0) {
                        // Head of stream - brightest
                        alpha = 1;
                        ctx.fillStyle = '#fff';
                    } else {
                        // Fade out towards tail
                        alpha = Math.max(0, 1 - (i / col.length) * 0.9);
                        const green = Math.floor(255 * alpha);
                        ctx.fillStyle = `rgb(0, ${green}, 0)`;
                    }

                    ctx.fillText(charData.char, col.x, y);
                }
            }
        },

        // ==================== Animation Loop ====================
        animate() {
            const theme = document.documentElement.getAttribute('data-theme');

            // Check if theme changed
            if (theme !== this.currentTheme) {
                this.currentTheme = theme;
                this.setupTheme();

                // Clear canvas on theme switch
                this.ctx.fillStyle = theme === 'cyber' ? '#0a0a0f' : '#000';
                this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height);
            }

            if (theme === 'cyber') {
                this.drawCyberBackground();
            } else {
                this.drawMatrixBackground();
            }

            this.animationId = requestAnimationFrame(() => this.animate());
        },

        destroy() {
            if (this.animationId) {
                cancelAnimationFrame(this.animationId);
            }
        }
    };

    // ==================== Main Application ====================

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

        // Initialize canvas background
        CanvasBackground.init();

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

            batchTable.appendChild(row);
        });

        // Bind cell copy events
        bindCellCopyEvents();
    }

    // Bind cell copy events
    function bindCellCopyEvents() {
        document.querySelectorAll('td.copyable').forEach(cell => {
            cell.addEventListener('click', () => {
                const value = cell.getAttribute('data-value');
                copyToClipboard(value);

                // Visual feedback
                cell.classList.add('copied');
                setTimeout(() => cell.classList.remove('copied'), 500);

                showToast('Copied');
            });
        });
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
