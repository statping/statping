context('Service Tests', () => {

    beforeEach(function () {
        cy.visit('http://localhost:8080/dashboard')
        cy.get('input[name="username"]').type('admin')
        cy.get('input[name="password"]').type('admin')
        cy.get('form').submit()
    })

    it('should view services', () => {
        cy.visit('http://localhost:8080/services')
        cy.get('tr').should('have.length', 6)
        cy.title().should('eq', 'Statup | Services')
    })

    it('should create HTTP GET service', () => {
        cy.visit('http://localhost:8080/services')
        cy.get('select[name="method"]').select('GET')
        cy.get('input[name="name"]').clear().type('Google.com')
        cy.get('select[name="check_type"]').select('http')
        cy.get('input[name="domain"]').clear().type('https://google.com')
        cy.get('input[name="expected_status"]').clear().type('200')
        cy.get('input[name="interval"]').clear().type('25')
        cy.get('input[name="timeout"]').clear().type('30')
        cy.get('form').submit()
        cy.title().should('eq', 'Statup | Services')
        cy.get('tr').should('have.length', 7)
    })

    it('should create HTTP POST service', () => {
        cy.visit('http://localhost:8080/services')
        cy.get('select[name="method"]').select('POST')
        cy.get('input[name="name"]').clear().type('JSON Regex Test')
        cy.get('select[name="check_type"]').select('http')
        cy.get('input[name="domain"]').clear().type('https://jsonplaceholder.typicode.com/posts')
        cy.get('textarea[name="post_data"]').clear().type(`(title)": "((\\"|[statup])*)"`)
        cy.get('input[name="expected_status"]').clear().type('201')
        cy.get('input[name="interval"]').clear().type('15')
        cy.get('input[name="timeout"]').clear().type('45')
        cy.get('form').submit()
        cy.title().should('eq', 'Statup | Services')
        cy.get('tr').should('have.length', 8)
    })

    it('should create TCP service', () => {
        cy.visit('http://localhost:8080/services')
        cy.get('select[name="check_type"]').select('tcp')
        cy.get('input[name="name"]').clear().type('Google DNS')
        cy.get('input[name="domain"]').clear().type('8.8.8.8')
        cy.get('input[name="port"]').clear().type('53')
        cy.get('input[name="interval"]').clear().type('25')
        cy.get('input[name="timeout"]').clear().type('15')
        cy.get('form').submit()
        cy.title().should('eq', 'Statup | Services')
        cy.get('tr').should('have.length', 9)
    })

    it('should view HTTP GET service', () => {
        cy.visit('http://localhost:8080/service/6')
        cy.title().should('eq', 'Statup | Google.com Service')
    })

    it('should view HTTP POST service', () => {
        cy.visit('http://localhost:8080/service/7')
        cy.title().should('eq', 'Statup | JSON Regex Test Service')
    })

    it('should view TCP service', () => {
        cy.visit('http://localhost:8080/service/8')
        cy.title().should('eq', 'Statup | Google DNS Service')
    })

    it('should update HTTP service', () => {
        cy.visit('http://localhost:8080/service/6')
        cy.title().should('eq', 'Statup | Google.com Service')
        cy.get('#service_name').clear().type('Google Updated')
        cy.get('#service_interval').clear().type('60')
        cy.get(':nth-child(3) > form').submit()
        cy.title().should('eq', 'Statup | Google Updated Service')
        cy.get('#service_name').should('have.value', 'Google Updated')
    });

    it('should check the updated service', () => {
        cy.visit('http://localhost:8080/service/6')
        cy.title().should('eq', 'Statup | Google Updated Service')
        cy.get('#service_name').should('have.value', 'Google Updated')
        cy.get('#service_interval').should('have.value', '60')
    })

    it('should delete a service', () => {
        cy.visit('http://localhost:8080/services')
        cy.get(':nth-child(5) > .text-right > .btn-group > .btn-danger').click()
        cy.title().should('eq', 'Statup | Services')
        cy.get('tr').should('have.length', 8)
    })


});