import React, { useEffect, useState } from 'react'

const Home = () => {
    const [message, setMessage] = useState(null)
    const fetchMessage = async () => {
        try {
            const res = await fetch("/api/message")
            if (!res.ok) {
                console.error("Error in response", res);
                return
            }
            const data = await res.json()
            setMessage(data.message)

        } catch (err) {
            console.error("Error caught: ", err);

        }
    }
    useEffect(() => {
        fetchMessage()
    }, [])
    return (
        <div>Home
            {message ? (
                <p>Message from server: {message}</p>
            ) : (
                <p>Waiting for message</p>
            )}
        </div>
    )
}

export default Home
