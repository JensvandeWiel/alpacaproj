import './bootstrap.ts'
import { createInertiaApp } from '@inertiajs/react'
import { createRoot } from 'react-dom/client'

createInertiaApp({
    resolve: name => {
        const pages = import.meta.glob('./Pages/**/*.tsx', { eager: true })
        return pages[`./Pages/${name}.tsx`]
    },
    setup({ el, App, props }) {
        const root = createRoot(el);
        if (App) {
            root.render(<App {...props} />);
        } else {
            console.error('Component not found or failed to load.');
        }
    }
})