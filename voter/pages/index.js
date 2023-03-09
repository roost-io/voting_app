import React, { Component } from 'react';
import Link from 'next/link'
import kubernates from '../public/assets/kubernates.png';

class Home extends Component {
	constructor(props) {
		super(props);
		this.handleonCardClick = this.handleonCardClick.bind(this);
		this.state = {
			candidates: [],
			candidate_id: '',
			voter_id: '',
			disabled: false,
			view: 1,
			showResultsButton: false,
			showNotification: false,
		};
	}

	componentDidMount() {
		let r = Math.random().toString(36).substring(7);
		this.setState({ voter_id: r });
		fetch(`/api/candidate`, {
			method: 'GET',
		})
			.then((response) => response.json())
			.then((response) => {
				console.log(response);

				this.setState({ candidates: response.Candidates });
			})
			.catch((error) => {
				console.error(
					'ballot service is not reachable at http://' + ec_server_endpoint
				);
			});
	}

	handleonCardClick(candidate) {
		if (this.state.disabled === false) {
			this.setState({ candidate_id: candidate.Name });
			this.setState({ disabled: true });

			const data = {
				candidate_id: this.state.candidate_id,
				vote: this.state.voter_id,
			};

			fetch(`/api/count`, {
				method: 'POST',
				body: JSON.stringify(data),
			})
				.then((response) => response.json())
				.then((response) => {
					if (response.success) {
						this.setState({ showResultsButton: true });
					}
				})
				.catch((error) => {
					console.error(
						'ballot service is not reachable at http://' + ballot_endpoint
					);
				});
		}
	}
	componentDidUpdate(prevProps, prevState) {
		// if (prevState.candidate_id !== this.state.candidate_id) {
		// 	const data = {
		// 		candidate_id: this.state.candidate_id,
		// 		vote: this.state.voter_id,
		// 	};

		// 	fetch(`/api/count`, {
		// 		method: 'POST',
		// 		body: JSON.stringify(data),
		// 	})
		// 		.then((response) => response.json())
		// 		.then((response) => {
		// 			if (response.success) {
		// 				this.setState({ showResultsButton: true });
		// 			}
		// 		})
		// 		.catch((error) => {
		// 			console.error(
		// 				'ballot service is not reachable at http://' + ballot_endpoint
		// 			);
		// 		});
		// }
		if (prevState.showResultsButton !== this.state.showResultsButton) {
			this.setState({ showNotification: true });
			this.timer = setTimeout(() => {
				this.setState({ showNotification: false });
			}, 3000);
		}
	}
	componentWillUnmount() {
		clearTimeout(this.timer);
	}

	render() {
		// const handleonCardClick = async (candidate) => {
		//   if (this.state.disabled === false) {
		// 		this.setState({ candidate_id: candidate.Name });
		// 		this.setState({ disabled: true });
		//   }
		// };
		// const showResults = () => {
		// 	this.props.history.push('/results');
		// };
		const CustomCard = (candidate, index) => {
			return (
				<div
					className={
						this.state.candidate_id === candidate.Name
							? 'card selectedCard'
							: 'card'
					}
					onClick={() => this.handleonCardClick(candidate)}
					key={index}
				>
					<div className="cardBackgroundContainer">
						<div className="cardBackground"></div>
						<div className="cardBackgroundImage">
							<img
								src={candidate.ImageUrl}
								width="150px"
								height="150px"
								className="image"
								alt={candidate.Name}
							/>
						</div>
					</div>
					<div className="cardContent">{candidate.Name}</div>
				</div>
			);
		};
		return (
			<div className="Home">
				<div className="logo">
					<img src={kubernates.src} width="70px" height="70px" alt={'logo'} />
				</div>
				<div className="heading">
					How do you create a K8S cluster on your local system ?
				</div>
				<div className="cardContainer">
					{this.state.candidates.map((candidate, index) => {
						return CustomCard(candidate, index);
					})}
				</div>
				{this.state.showResultsButton && (
					<Link className="showResultsButton" href='/results'>
						Show Results
					</Link>
				)}
				{this.state.showNotification && (
					<div className="notificationPopup">
						<div
							className="closeNotification"
							onClick={() => this.setState({ showNotification: false })}
						>
							x
						</div>
						<div className="notificationContent">
							Vote registered for {this.state.candidate_id}
						</div>
					</div>
				)}
			</div>
		);
	}
}

export default Home;
