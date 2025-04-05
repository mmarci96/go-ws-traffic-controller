import { useState } from 'react'

const Form = () => {
    const [name, setName] = useState("")
    const handleSubmit = async (e) => {
        e.preventDefault()
        try {
            const res = await fetch("/api/name", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ name })
            })
            if (!res.ok) {
                console.error("Error from response: ", res);
                return;
            }
            const data = await res.json()
            console.log(data)
        } catch (err) {
            console.error("Caught error: ", e);
        }
    }

    return (
        <div>
            Form
            <form onSubmit={handleSubmit}>
                <input
                    type='text'
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                />
                <button type='submit'>
                    submit
                </button>
            </form>
        </div>
    )
}

export default Form
