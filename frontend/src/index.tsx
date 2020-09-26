import * as React from "react";
import * as ReactDOM from "react-dom";
import {HashRouter as Router, Route} from "react-router-dom";
import {EventListContainer} from "./components/event_list_container";
import {Navigation} from "./components/navigation";
import {EventBookingFormContainer} from "./components/event_booking_form_container";

import 'bootstrap/dist/css/bootstrap.min.css';

import { events,booking} from "./config"

class App extends React.Component<{}, {}> {
    render() {
        const eventList = () => <EventListContainer eventServiceURL={events}/>;

        const eventBooking = ({ match }: any) => <EventBookingFormContainer eventID={match.params.id} eventServiceURL={events} bookingServiceURL={booking}/>;

        return <Router>
            <div>
                <Navigation brandName="MyEvents"/>
                <div className="container">
                    <h1>My Events</h1>

                    <Route exact path="/" component={eventList}/>
                    <Route path="/events/:id/book" component={eventBooking}/>
                </div>
            </div>
        </Router>
    }
}

ReactDOM.render(
    <App/>,
    document.getElementById("myevents-app")
);