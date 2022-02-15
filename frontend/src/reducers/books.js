const bookReducer = (books = [], action) => {
    switch (action.type) {
        case 'TOP':
            return [...action?.data]
        case 'LIKE':
            const index = books.findIndex((book) => book.id === action?.ids.story)
            const newBooks = books
            if (newBooks[index].likes.includes(action.ids.id)) {
                newBooks[index].likes = newBooks[index].likes.filter((bookid) => { return bookid !== action.ids.id })
            } else {
                newBooks[index].likes.push(action.ids.id)
            }
            return [...books]
        default:
            return books;
    }

}

export default bookReducer