import React from 'react'
import Moment from 'react-moment';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';
import FavoriteIcon from '@mui/icons-material/Favorite';
import { useSelector, useDispatch } from 'react-redux';
import { pink } from '@mui/material/colors';
import IconButton from '@mui/material/IconButton';
import { likeBook } from "../../actions/books"
import { styled } from '@mui/material/styles';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import Collapse from '@mui/material/Collapse'
import CardHeader from '@mui/material/CardHeader';
const Story = ({ data }) => {
    const id = useSelector((state) => state.Auth.id)
    const dispatch = useDispatch();
    const handleExpandClick = () => {
        setExpanded(!expanded);
    };

    const ExpandMore = styled((props) => {
        const { expand, ...other } = props;
        return <IconButton {...other} />;
    })(({ theme, expand }) => ({
        transform: !expand ? 'rotate(0deg)' : 'rotate(180deg)',
        marginLeft: 'auto',
        transition: theme.transitions.create('transform', {
            duration: theme.transitions.duration.shortest,
        }),
    }));
    const [expanded, setExpanded] = React.useState(false);
    return (
        <div>
            <Card sx={{ maxWidth: 345, minWidth: 345, margin: "auto" }}>
                <CardHeader
                    title={data.title}
                    subheader={<Moment fromNow>{data.published_date}</Moment>}
                />
                <CardActions disableSpacing>
                    <IconButton onClick={() => { dispatch(likeBook({ id, story: data.id })) }}>
                        {data.likes.includes(id) ? (<FavoriteIcon sx={{ color: pink[500] }} />) : (<FavoriteBorderIcon />)}  {data.likes.length}
                    </IconButton>
                    <ExpandMore
                        expand={expanded}
                        onClick={handleExpandClick}
                        aria-expanded={expanded}
                        aria-label="show more"
                    >
                        <ExpandMoreIcon />
                    </ExpandMore>
                </CardActions>
                <Collapse in={expanded} timeout="auto" unmountOnExit>
                    <CardContent>
                        <Typography paragraph>{data.story}</Typography>

                    </CardContent>
                </Collapse>
            </Card>
        </div>
    )
}

export default Story