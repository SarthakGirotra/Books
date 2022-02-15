import React, { useLayoutEffect } from 'react'
import { useDispatch, useSelector } from "react-redux"
import { logout } from "../../actions/user"
import { getTopBooks } from "../../actions/books"
import Story from '../story/story'
import { Button } from '@mui/material'

function Homepage() {
    var books
    const dispatch = useDispatch();
    /*eslint-disable */
    useLayoutEffect(() => {
        dispatch(getTopBooks())
    }, [])
    books = useSelector((state) => state.books)
    /*eslint-enable */

    return (
        <div>
            <div className='container'>
                <Button variant='contained' sx={{ margin: "10px 30px 0 30px", maxWidth: "100px" }} onClick={() => { dispatch(logout()) }}>Logout</Button>
                <div className='books'>
                    {books.map((data) => <Story data={data} key={data.id} />)}
                </div>
            </div>
        </div>
    )
}
export default Homepage      