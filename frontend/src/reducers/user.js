const authReducer = (state = { id: null }, action) => {
    switch (action.type) {
        case 'AUTH':
            localStorage.setItem("id", action?.data?.id)
            return { ...state, id: action?.data?.id }
        case 'LOGOUT':
            localStorage.removeItem('id')
            return { ...state, id: null }
        case 'HOME':
            return { ...state, id: action?.payload }
        default:
            return state
    }
}

export default authReducer;