import React from 'react'
import { createRoot } from 'react-dom/client'
import './style.css'
import App from './App'
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <React.StrictMode>
        <App />
        <ToastContainer />
    </React.StrictMode>
)
