import { BrowserRouter, Route, Routes } from 'react-router';
import './App.css'
import Form from './scenes/Form'
import Home from './scenes/Home';

function App() {

    return (

        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/form" element={<Form />} />
            </Routes>
        </BrowserRouter>
    )
}

export default App
