const errorReducer = (error = null, action) => {
    switch (action.type) {
        case "ERROR":
            return error = action?.data
        case "REMOVE_ERROR":
            return error = null
        default:
            return error
    }
}
export default errorReducer